package api

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	yaml "gopkg.in/yaml.v2"
)

type DeployedProductOutput struct {
	Type string
	GUID string
}

func (a Api) GetDeployedProductManifest(guid string) (string, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("/api/v0/deployed/products/%s/manifest", guid), nil)
	if err != nil {
		return "", err
	}

	resp, err := a.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("could not make api request to staged products manifest endpoint: %s", err)
	}

	if err = validateStatusOK(resp); err != nil {
		return "", err
	}

	defer resp.Body.Close()
	var contents interface{}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if err = yaml.Unmarshal(body, &contents); err != nil {
		return "", fmt.Errorf("could not parse json: %s", err)
	}

	manifest, err := yaml.Marshal(contents)
	if err != nil {
		return "", err // this should never happen, all valid json can be marshalled
	}

	return string(manifest), nil
}

func (a Api) ListDeployedProducts() ([]DeployedProductOutput, error) {
	req, err := http.NewRequest("GET", "/api/v0/deployed/products", nil)
	if err != nil {
		return []DeployedProductOutput{}, err
	}

	resp, err := a.client.Do(req)
	if err != nil {
		return []DeployedProductOutput{}, fmt.Errorf("could not make api request to deployed products endpoint: %s", err)
	}
	defer resp.Body.Close()

	if err = validateStatusOK(resp); err != nil {
		return []DeployedProductOutput{}, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []DeployedProductOutput{}, err
	}

	var deployedProducts []DeployedProductOutput
	err = json.Unmarshal(respBody, &deployedProducts)
	if err != nil {
		return []DeployedProductOutput{}, fmt.Errorf("could not unmarshal deployed products response: %s", err)
	}

	return deployedProducts, nil
}

func (a Api) GetDeployedProductProperties(product string) (map[string]ResponseProperty, error) {
	respBody, err := a.fetchDeployedProductResource(product, "properties")
	if err != nil {
		return nil, err
	}
	defer respBody.Close()

	propertiesResponse := struct {
		Properties map[string]ResponseProperty `json:"properties"`
	}{}
	if err = json.NewDecoder(respBody).Decode(&propertiesResponse); err != nil {
		return nil, fmt.Errorf("could not parse json: %s", err)
	}

	return propertiesResponse.Properties, nil
}

func (a Api) GetDeployedProductNetworksAndAZs(product string) (map[string]interface{}, error) {
	respBody, err := a.fetchDeployedProductResource(product, "networks_and_azs")
	if err != nil {
		return nil, err
	}
	defer respBody.Close()

	networksResponse := struct {
		Networks map[string]interface{} `json:"networks_and_azs"`
	}{}
	if err = json.NewDecoder(respBody).Decode(&networksResponse); err != nil {
		return nil, fmt.Errorf("could not parse json: %s", err)
	}

	return networksResponse.Networks, nil
}

func (a Api) fetchDeployedProductResource(guid, endpoint string) (io.ReadCloser, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("/api/v0/deployed/products/%s/%s", guid, endpoint), nil)
	if err != nil {
		return nil, err // un-tested
	}

	resp, err := a.client.Do(req)
	if err != nil {
		return nil,
			fmt.Errorf("could not make api request to deployed product properties endpoint: %s", err)
	}

	if err = validateStatusOK(resp); err != nil {
		return nil, err
	}

	return resp.Body, nil
}
