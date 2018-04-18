package commands

import (
	"fmt"

	"github.com/pivotal-cf/jhanda"
	"github.com/pivotal-cf/om/api"
	yaml "gopkg.in/yaml.v2"
)

type DeployedConfig struct {
	logger  logger
	service deployedConfigService
	Options struct {
		Product string `long:"product-name"    short:"p" required:"true" description:"name of product"`
	}
}

//go:generate counterfeiter -o ./fakes/deployed_config_service.go --fake-name DeployedConfigService . deployedConfigService
type deployedConfigService interface {
	GetDeployedProductByName(product string) (api.DeployedProductsFindOutput, error)
	ListDeployedProductJobs(productGUID string) (map[string]string, error)
	GetDeployedProductJobResourceConfig(productGUID, jobGUID string) (api.JobProperties, error)
	GetDeployedProductProperties(product string) (map[string]api.ResponseProperty, error)
	GetDeployedProductNetworksAndAZs(product string) (map[string]interface{}, error)
	GetDeployedProductCredential(product string, credentialName string) (api.CredentialOutput, error)
}

func NewDeployedConfig(service deployedConfigService, logger logger) DeployedConfig {
	return DeployedConfig{
		logger:  logger,
		service: service,
	}
}

func (ec DeployedConfig) Usage() jhanda.Usage {
	return jhanda.Usage{
		Description:      "This command generates a config from a deployed product that can be passed in to om configure-product (Note: this file will contain your credentials in plain-text)",
		ShortDescription: "generates a config from a deployed product",
		Flags:            ec.Options,
	}
}

func (ec DeployedConfig) Execute(args []string) error {
	if _, err := jhanda.Parse(&ec.Options, args); err != nil {
		return fmt.Errorf("could not parse deployed-config flags: %s", err)
	}

	findOutput, err := ec.service.GetDeployedProductByName(ec.Options.Product)
	if err != nil {
		return err
	}
	productGUID := findOutput.Product.GUID

	properties, err := ec.service.GetDeployedProductProperties(productGUID)
	if err != nil {
		return err
	}

	configurableProperties := map[string]interface{}{}

	for name, property := range properties {
		if property.Configurable && property.Value != nil {
			if property.IsCredential {
				credResult, err := ec.service.GetDeployedProductCredential(productGUID, name)
				if err != nil {
					return err
				}
				configurableProperties[name] = map[string]interface{}{"value": credResult.Credential.Value}
			} else {
				configurableProperties[name] = map[string]interface{}{"value": property.Value}
			}
		}
	}

	networks, err := ec.service.GetDeployedProductNetworksAndAZs(productGUID)
	if err != nil {
		return err
	}

	jobs, err := ec.service.ListDeployedProductJobs(productGUID)
	if err != nil {
		return err
	}

	resourceConfig := map[string]api.JobProperties{}

	for name, jobGUID := range jobs {
		jobProperties, err := ec.service.GetDeployedProductJobResourceConfig(productGUID, jobGUID)
		if err != nil {
			return err
		}

		resourceConfig[name] = jobProperties
	}

	config := struct {
		Properties               map[string]interface{}       `yaml:"product-properties"`
		NetworkProperties        map[string]interface{}       `yaml:"network-properties"`
		ResourceConfigProperties map[string]api.JobProperties `yaml:"resource-config"`
	}{
		Properties:               configurableProperties,
		NetworkProperties:        networks,
		ResourceConfigProperties: resourceConfig,
	}

	output, err := yaml.Marshal(config)
	if err != nil {
		return fmt.Errorf("failed to unmarshal config: %s", err) // un-tested
	}
	ec.logger.Println(string(output))

	return nil
}
