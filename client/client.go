package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/terraform-provider-aviatrix-cloud-poc/config"
)

type HttpHandler interface {
	Request(ctx context.Context, method string, url string, i interface{}, output interface{}) (httpStatusCode int, err error)
	Get(ctx context.Context, url string, i interface{}, output interface{}) (httpStatusCode int, err error)
	Post(ctx context.Context, url string, i interface{}, output interface{}) (httpStatusCode int, err error)
}

type Client struct {
	Host       string
	Username   string
	Password   string
	HTTPClient *http.Client
}

func NewClient(host, username, password string) (HttpHandler, error) {
	var err error
	client := &Client{
		Host:       host,
		Username:   username,
		Password:   password,
		HTTPClient: &http.Client{},
	}

	return client, err
}

// Get issues an HTTP GET request
func (c *Client) Get(ctx context.Context, path string, payload interface{}, output interface{}) (rhttpStatusCode int, err error) {
	return c.Request(ctx, http.MethodGet, path, payload, output)
}

// Post issues an HTTP POST request
func (c *Client) Post(ctx context.Context, path string, payload interface{}, output interface{}) (httpStatusCode int, err error) {
	return c.Request(ctx, http.MethodPost, path, payload, output)
}

func (c *Client) Request(ctx context.Context, method string, path string, payload interface{}, output interface{}) (httpStatusCode int, err error) {
	url := c.Host + path
	debugMsg := fmt.Sprintf(" ====> method: %v, url: %v, payload : %v ", method, url, payload)
	//panic(fmt.Sprintf("======> Full URL : %v ", url))
	tflog.Debug(ctx, debugMsg)
	// Marshal the interface{} to JSON (byte slice)
	byteData, err := json.Marshal(payload)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Error marshaling method: %v, url :%v, error: %v", method, url, err.Error()))
		return http.StatusBadRequest, err
	}

	httpResponse, err := RetryableHttpRequest(method, url, byteData)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Error response  for method: %v, url :%v, error: %v", method, url, err.Error()))
		return http.StatusInternalServerError, err
	}

	defer httpResponse.Body.Close()
	byteResp, err := io.ReadAll(httpResponse.Body)

	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Error while reading response  for method: %v, url :%v, error: %v", method, url, err.Error()))
		return http.StatusInternalServerError, err
	}

	tflog.Debug(ctx, fmt.Sprintf("===> Response : %v ", string(byteResp)))

	// Unmarshal the JSON into the struct
	err = json.Unmarshal(byteResp, output)
	if err != nil {
		//errMsg := fmt.Sprintf("Error unmarshalling JSON: %v", err.Error())
		return http.StatusInternalServerError, err
	}

	return httpResponse.StatusCode, nil
}

func RetryableHttpRequest(httpMethod, url string, data []byte) (*http.Response, error) {

	var req *http.Request
	var err error

	// Create the request (either GET or POST)
	if httpMethod == http.MethodPost {
		req, err = http.NewRequest(httpMethod, url, bytes.NewBuffer(data))
	} else {
		req, err = http.NewRequest(httpMethod, url, nil)
	}
	
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	retryClient := retryablehttp.NewClient()
	retryClient.RetryWaitMin = config.HttpRetryWaitMin
	retryClient.RetryWaitMax = config.HttpRetryWaitMax
	retryClient.RetryMax = config.HttpRetryMax

	standardClient := retryClient.StandardClient()

	resp, err := standardClient.Do(req)
	if err == nil && resp.StatusCode < 500 { // If request is successful or a 4xx error (non-retryable)
		return resp, nil
	}

	return nil, fmt.Errorf("request failed after %d retries, url : %s, method : %s ", config.HttpRetryMax, url, httpMethod)

}
