package commands_test

import (
	"errors"

	"github.com/pivotal-cf/jhanda"
	"github.com/pivotal-cf/om/api"
	"github.com/pivotal-cf/om/commands"
	"github.com/pivotal-cf/om/commands/fakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DeployedConfig", func() {
	var (
		logger      *fakes.Logger
		fakeService *fakes.DeployedConfigService
	)

	BeforeEach(func() {
		logger = &fakes.Logger{}

		fakeService = &fakes.DeployedConfigService{}
		fakeService.GetDeployedProductPropertiesReturns(
			map[string]api.ResponseProperty{
				".properties.some-string-property": api.ResponseProperty{
					Value:        "some-value",
					Configurable: true,
				},
				".properties.some-non-configurable-property": api.ResponseProperty{
					Value:        "some-value",
					Configurable: false,
				},
				".properties.some-secret-property": api.ResponseProperty{
					Value: map[string]interface{}{
						"secret": "***",
					},
					IsCredential: true,
					Configurable: true,
				},
				".properties.some-null-property": api.ResponseProperty{
					Value:        nil,
					Configurable: true,
				},
			}, nil)
		fakeService.GetDeployedProductNetworksAndAZsReturns(
			map[string]interface{}{
				"singleton_availability_zone": map[string]string{
					"name": "az-one",
				},
			}, nil)

		fakeService.GetDeployedProductByNameReturns(api.DeployedProductsFindOutput{
			Product: api.DeployedProductOutput{
				GUID: "some-product-guid",
			},
		}, nil)

		fakeService.ListDeployedProductJobsReturns(map[string]string{
			"some-job": "some-job-guid",
		}, nil)
		fakeService.GetDeployedProductJobResourceConfigReturns(api.JobProperties{
			InstanceType: api.InstanceType{
				ID: "automatic",
			},
			Instances: 1,
		}, nil)
		fakeService.GetDeployedProductCredentialReturns(api.CredentialOutput{
			Credential: api.Credential{
				Type: "some-secret-type",
				Value: map[string]string{
					"secret": "some-secret-value",
				},
			},
		}, nil)
	})

	Describe("Execute", func() {
		It("writes a config file to output", func() {
			command := commands.NewDeployedConfig(fakeService, logger)
			err := command.Execute([]string{
				"--product-name", "some-product",
			})
			Expect(err).NotTo(HaveOccurred())

			Expect(fakeService.GetDeployedProductByNameCallCount()).To(Equal(1))
			Expect(fakeService.GetDeployedProductByNameArgsForCall(0)).To(Equal("some-product"))

			Expect(fakeService.GetDeployedProductPropertiesCallCount()).To(Equal(1))
			Expect(fakeService.GetDeployedProductPropertiesArgsForCall(0)).To(Equal("some-product-guid"))

			Expect(fakeService.GetDeployedProductNetworksAndAZsCallCount()).To(Equal(1))
			Expect(fakeService.GetDeployedProductNetworksAndAZsArgsForCall(0)).To(Equal("some-product-guid"))

			Expect(fakeService.ListDeployedProductJobsCallCount()).To(Equal(1))
			Expect(fakeService.ListDeployedProductJobsArgsForCall(0)).To(Equal("some-product-guid"))

			Expect(fakeService.GetDeployedProductJobResourceConfigCallCount()).To(Equal(1))
			productGuid, jobsGuid := fakeService.GetDeployedProductJobResourceConfigArgsForCall(0)
			Expect(productGuid).To(Equal("some-product-guid"))
			Expect(jobsGuid).To(Equal("some-job-guid"))

			Expect(fakeService.GetDeployedProductCredentialCallCount()).To(Equal(1))
			productGuid, credName := fakeService.GetDeployedProductCredentialArgsForCall(0)
			Expect(productGuid).To(Equal("some-product-guid"))
			Expect(credName).To(Equal(".properties.some-secret-property"))

			Expect(logger.PrintlnCallCount()).To(Equal(1))
			output := logger.PrintlnArgsForCall(0)
			Expect(output).To(ContainElement(MatchYAML(`---
product-properties:
  .properties.some-string-property:
    value: some-value
  .properties.some-secret-property:
    value:
      secret: some-secret-value
network-properties:
  singleton_availability_zone:
    name: az-one
resource-config:
  some-job:
    instances: 1
    instance_type:
      id: automatic
`)))
		})
	})

	Context("failure cases", func() {
		Context("when an unknown flag is provided", func() {
			It("returns an error", func() {
				command := commands.NewDeployedConfig(fakeService, logger)
				err := command.Execute([]string{"--badflag"})
				Expect(err).To(MatchError("could not parse deployed-config flags: flag provided but not defined: -badflag"))
			})
		})

		Context("when product name is not provided", func() {
			It("returns an error and prints out usage", func() {
				command := commands.NewDeployedConfig(fakeService, logger)
				err := command.Execute([]string{})
				Expect(err).To(MatchError("could not parse deployed-config flags: missing required flag \"--product-name\""))
			})
		})

		Context("when looking up the product GUID fails", func() {
			BeforeEach(func() {
				fakeService.GetDeployedProductByNameReturns(api.DeployedProductsFindOutput{}, errors.New("some-error"))
			})

			It("returns an error", func() {
				command := commands.NewDeployedConfig(fakeService, logger)
				err := command.Execute([]string{
					"--product-name", "some-product",
				})
				Expect(err).To(MatchError("some-error"))
			})
		})

		Context("when looking up the product properties fails", func() {
			BeforeEach(func() {
				fakeService.GetDeployedProductPropertiesReturns(nil, errors.New("some-error"))
			})

			It("returns an error", func() {
				command := commands.NewDeployedConfig(fakeService, logger)
				err := command.Execute([]string{
					"--product-name", "some-product",
				})
				Expect(err).To(MatchError("some-error"))
			})
		})

		Context("when looking up the network fails", func() {
			BeforeEach(func() {
				fakeService.GetDeployedProductNetworksAndAZsReturns(nil, errors.New("some-error"))
			})

			It("returns an error", func() {
				command := commands.NewDeployedConfig(fakeService, logger)
				err := command.Execute([]string{
					"--product-name", "some-product",
				})
				Expect(err).To(MatchError("some-error"))
			})
		})

		Context("when listing jobs fails", func() {
			BeforeEach(func() {
				fakeService.ListDeployedProductJobsReturns(nil, errors.New("some-error"))
			})

			It("returns an error", func() {
				command := commands.NewDeployedConfig(fakeService, logger)
				err := command.Execute([]string{
					"--product-name", "some-product",
				})
				Expect(err).To(MatchError("some-error"))
			})
		})

		Context("when looking up the job fails", func() {
			BeforeEach(func() {
				fakeService.GetDeployedProductJobResourceConfigReturns(api.JobProperties{}, errors.New("some-error"))
			})

			It("returns an error", func() {
				command := commands.NewDeployedConfig(fakeService, logger)
				err := command.Execute([]string{
					"--product-name", "some-product",
				})
				Expect(err).To(MatchError("some-error"))
			})
		})

		Context("when looking up the credential fails", func() {
			BeforeEach(func() {
				fakeService.GetDeployedProductCredentialReturns(api.CredentialOutput{}, errors.New("some-error"))
			})

			It("returns an error", func() {
				command := commands.NewDeployedConfig(fakeService, logger)
				err := command.Execute([]string{
					"--product-name", "some-product",
				})
				Expect(err).To(MatchError("some-error"))
			})
		})
	})

	Describe("Usage", func() {
		It("returns usage information for the command", func() {
			command := commands.NewDeployedConfig(nil, nil)
			Expect(command.Usage()).To(Equal(jhanda.Usage{
				Description:      "This command generates a config from a deployed product that can be passed in to om configure-product (Note: this file will contain your credentials in plain-text)",
				ShortDescription: "generates a config from a deployed product",
				Flags:            command.Options,
			}))
		})
	})
})
