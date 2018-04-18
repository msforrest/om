package api_test

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/pivotal-cf/om/api"
	"github.com/pivotal-cf/om/api/fakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DeployedProducts", func() {
	var (
		client  *fakes.HttpClient
		service api.Api
	)

	BeforeEach(func() {
		client = &fakes.HttpClient{}
		service = api.New(api.ApiInput{
			Client: client,
		})
	})

	Describe("GetDeployedProductManifest", func() {
		BeforeEach(func() {
			client.DoStub = func(req *http.Request) (*http.Response, error) {
				var resp *http.Response
				switch req.URL.Path {
				case "/api/v0/deployed/products/some-product-guid/manifest":
					resp = &http.Response{
						StatusCode: http.StatusOK,
						Body: ioutil.NopCloser(bytes.NewBufferString(`{
							"key-1": {
								"key-2": "value-1"
							},
							"key-3": "value-2",
							"key-4": 2147483648
						}`)),
					}
				}
				return resp, nil
			}
		})

		It("returns a manifest of a product", func() {
			manifest, err := service.GetDeployedProductManifest("some-product-guid")
			Expect(err).NotTo(HaveOccurred())
			Expect(manifest).To(MatchYAML(`---
key-1:
  key-2: value-1
key-3: value-2
key-4: 2147483648
`))
		})

		Context("failure cases", func() {
			Context("when the request object is invalid", func() {
				It("returns an error", func() {
					_, err := service.GetDeployedProductManifest("invalid-guid-%%%")
					Expect(err).To(MatchError(ContainSubstring("invalid URL escape")))
				})
			})

			Context("when the client request fails", func() {
				It("returns an error", func() {
					client.DoReturns(&http.Response{}, errors.New("nope"))

					_, err := service.GetDeployedProductManifest("some-product-guid")
					Expect(err).To(MatchError("could not make api request to staged products manifest endpoint: nope"))
				})
			})

			Context("when the server returns a non-200 status code", func() {
				It("returns an error", func() {
					client.DoReturns(&http.Response{
						StatusCode: http.StatusTeapot,
						Body:       ioutil.NopCloser(bytes.NewBufferString("")),
					}, nil)

					_, err := service.GetDeployedProductManifest("some-product-guid")
					Expect(err).To(MatchError(ContainSubstring("request failed: unexpected response")))
				})
			})

			Context("when the returned JSON is invalid", func() {
				It("returns an error", func() {
					client.DoReturns(&http.Response{
						StatusCode: http.StatusOK,
						Body:       ioutil.NopCloser(bytes.NewBufferString("%%%")),
					}, nil)

					_, err := service.GetDeployedProductManifest("some-product-guid")
					Expect(err).To(MatchError(ContainSubstring("could not parse json")))
				})
			})
		})
	})

	Describe("List", func() {
		BeforeEach(func() {
			client.DoStub = func(req *http.Request) (*http.Response, error) {
				var resp *http.Response
				resp = &http.Response{
					StatusCode: http.StatusOK,
					Body:       ioutil.NopCloser(bytes.NewBufferString(``)),
				}
				switch req.URL.Path {
				case "/api/v0/deployed/products":
					resp = &http.Response{
						StatusCode: http.StatusOK,
						Body: ioutil.NopCloser(bytes.NewBufferString(`[{
							"guid":"some-product-guid",
							"type":"some-type"
						},
						{
							"guid":"some-other-product-guid",
							"type":"some-other-type"
						}]`)),
					}
				}
				return resp, nil
			}
		})

		It("retrieves a list of deployed products from the Ops Manager", func() {
			output, err := service.ListDeployedProducts()
			Expect(err).NotTo(HaveOccurred())

			Expect(output).To(Equal([]api.DeployedProductOutput{
				{
					GUID: "some-product-guid",
					Type: "some-type",
				},
				{
					GUID: "some-other-product-guid",
					Type: "some-other-type",
				},
			},
			))

			Expect(client.DoCallCount()).To(Equal(1))

			By("checking for deployed products")
			avReq := client.DoArgsForCall(0)
			Expect(avReq.URL.Path).To(Equal("/api/v0/deployed/products"))
		})

		Context("failure cases", func() {
			Context("when the request fails", func() {
				BeforeEach(func() {
					client.DoReturns(&http.Response{}, errors.New("nope"))
				})

				It("returns an error", func() {
					_, err := service.ListDeployedProducts()
					Expect(err).To(MatchError("could not make api request to deployed products endpoint: nope"))
				})
			})

			Context("when the server returns a non-200 status code", func() {
				BeforeEach(func() {
					client.DoReturns(&http.Response{
						StatusCode: http.StatusTeapot,
						Body:       ioutil.NopCloser(bytes.NewBufferString("")),
					}, nil)
				})

				It("returns an error", func() {
					_, err := service.ListDeployedProducts()
					Expect(err).To(MatchError(ContainSubstring("request failed: unexpected response")))
				})
			})

			Context("when the server returns invalid JSON", func() {
				BeforeEach(func() {
					client.DoReturns(&http.Response{
						StatusCode: http.StatusOK,
						Body:       ioutil.NopCloser(bytes.NewBufferString("%%")),
					}, nil)
				})

				It("returns an error", func() {
					_, err := service.ListDeployedProducts()
					Expect(err).To(MatchError(ContainSubstring("could not unmarshal deployed products response:")))
				})
			})
		})
	})

	Describe("GetDeployedProductProperties", func() {
		BeforeEach(func() {
			client.DoStub = func(req *http.Request) (*http.Response, error) {
				var resp *http.Response
				switch req.URL.Path {
				case "/api/v0/deployed/products/some-product-guid/properties":
					resp = &http.Response{
						StatusCode: http.StatusOK,
						Body: ioutil.NopCloser(bytes.NewBufferString(`{
							"properties": {
								".properties.some-configurable-property": {
									"value": "some-value",
									"configurable": true
								},
								".properties.some-non-configurable-property": {
									"value": "some-value",
									"configurable": false
								},
								".properties.some-secret-property": {
									"value": {
										"some-secret-type": "***"
									},
									"configurable": true,
									"credential": true
								}
							}
						}`)),
					}
				}
				return resp, nil
			}
		})

		It("returns the configuration for a product", func() {
			config, err := service.GetDeployedProductProperties("some-product-guid")
			Expect(err).NotTo(HaveOccurred())
			Expect(config).To(Equal(map[string]api.ResponseProperty{
				".properties.some-configurable-property": api.ResponseProperty{
					Value:        "some-value",
					Configurable: true,
				},
				".properties.some-non-configurable-property": api.ResponseProperty{
					Value:        "some-value",
					Configurable: false,
				},
				".properties.some-secret-property": api.ResponseProperty{
					Value: map[string]interface{}{
						"some-secret-type": "***",
					},
					Configurable: true,
					IsCredential: true,
				},
			}))
		})

		Context("failure cases", func() {
			Context("when the properties request returns an error", func() {
				BeforeEach(func() {
					client.DoStub = func(req *http.Request) (*http.Response, error) {
						var resp *http.Response
						switch req.URL.Path {
						case "/api/v0/deployed/products/some-product-guid/properties":
							return &http.Response{}, errors.New("some-error")
						}
						return resp, nil
					}
				})
				It("returns an error", func() {
					_, err := service.GetDeployedProductProperties("some-product-guid")
					Expect(err).To(MatchError(`could not make api request to deployed product properties endpoint: some-error`))
				})
			})

			Context("when the properties request returns a non 200 error code", func() {
				BeforeEach(func() {
					client.DoStub = func(req *http.Request) (*http.Response, error) {
						var resp *http.Response
						switch req.URL.Path {
						case "/api/v0/deployed/products/some-product-guid/properties":
							return &http.Response{
								StatusCode: http.StatusTeapot,
								Body:       ioutil.NopCloser(bytes.NewBufferString("")),
							}, nil
						}
						return resp, nil
					}
				})
				It("returns an error", func() {
					_, err := service.GetDeployedProductProperties("some-product-guid")
					Expect(err).To(MatchError(ContainSubstring("request failed: unexpected response")))
				})
			})

			Context("when the server returns invalid json", func() {
				BeforeEach(func() {
					client.DoStub = func(req *http.Request) (*http.Response, error) {
						var resp *http.Response
						switch req.URL.Path {
						case "/api/v0/deployed/products/some-product-guid/properties":
							resp = &http.Response{
								StatusCode: http.StatusOK,
								Body:       ioutil.NopCloser(bytes.NewBufferString(`{{{`)),
							}
						}
						return resp, nil
					}
				})

				It("returns an error", func() {
					_, err := service.GetDeployedProductProperties("some-product-guid")
					Expect(err).To(MatchError(ContainSubstring("could not parse json")))
				})
			})
		})
	})

	Describe("GetDeployedProductNetworksAndAZs", func() {
		BeforeEach(func() {
			client.DoStub = func(req *http.Request) (*http.Response, error) {
				var resp *http.Response
				switch req.URL.Path {
				case "/api/v0/deployed/products/some-product-guid/networks_and_azs":
					resp = &http.Response{
						StatusCode: http.StatusOK,
						Body: ioutil.NopCloser(bytes.NewBufferString(`{
							"networks_and_azs": {
						  	"singleton_availability_zone": {
                  "name": "az-one"
                },
                "other_availability_zones": [
                  { "name": "az-two" },
                  { "name": "az-three" }
                ],
                "network": {
                  "name": "network-one"
                }
						  }
						}`)),
					}
				}
				return resp, nil
			}
		})

		It("returns the networks + azs for a product", func() {
			config, err := service.GetDeployedProductNetworksAndAZs("some-product-guid")
			Expect(err).NotTo(HaveOccurred())
			Expect(config).To(Equal(map[string]interface{}{
				"singleton_availability_zone": map[string]interface{}{
					"name": "az-one",
				},
				"other_availability_zones": []interface{}{
					map[string]interface{}{"name": "az-two"},
					map[string]interface{}{"name": "az-three"},
				},
				"network": map[string]interface{}{
					"name": "network-one",
				},
			}))
		})

		Context("failure cases", func() {
			Context("when the networks_and_azs request returns an error", func() {
				BeforeEach(func() {
					client.DoStub = func(req *http.Request) (*http.Response, error) {
						var resp *http.Response
						switch req.URL.Path {
						case "/api/v0/deployed/products/some-product-guid/networks_and_azs":
							return &http.Response{}, errors.New("some-error")
						}
						return resp, nil
					}
				})

				It("returns an error", func() {
					_, err := service.GetDeployedProductNetworksAndAZs("some-product-guid")
					Expect(err).To(MatchError(`could not make api request to deployed product properties endpoint: some-error`))
				})
			})

			Context("when the server returns invalid json", func() {
				BeforeEach(func() {
					client.DoStub = func(req *http.Request) (*http.Response, error) {
						var resp *http.Response
						switch req.URL.Path {
						case "/api/v0/deployed/products/some-product-guid/networks_and_azs":
							resp = &http.Response{
								StatusCode: http.StatusOK,
								Body:       ioutil.NopCloser(bytes.NewBufferString(`{{{`)),
							}
						}
						return resp, nil
					}
				})

				It("returns an error", func() {
					_, err := service.GetDeployedProductNetworksAndAZs("some-product-guid")
					Expect(err).To(MatchError(ContainSubstring("could not parse json")))
				})
			})
		})
	})
})
