/*
Copyright © 2022 Richard Wooding richard.wooding@gmail.com

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package api

import (
	"bytes"
	"encoding/json"
	"github.com/richardwooding/powerdown/model"
	"io"
	"net/http"
	"net/url"
	"time"
)

type RestClient struct {
	baseUrl *url.URL
	userAgent string
	token string
	httpClient http.Client
}

func NewRestClient(token string, timeout time.Duration) (*RestClient, error){
	restClient := new(RestClient)
	restClient.userAgent = "powerdown/0.0.1 https://github.com/richardwooding/powerdown"
	restClient.token = token
	restClient.httpClient = http.Client{
		Timeout: timeout,
	}
	baseUrl, err := url.Parse("https://developer.sepush.co.za/business/2.0/")
	if (err == nil) {
		restClient.baseUrl = baseUrl
	}
	return restClient, err
}



func (c *RestClient) Allowance() (*model.AllowanceResponse, error) {
	req, err := c.newRequest(http.MethodGet, "./api_allowance", nil)
	if err != nil {
		return nil, err
	}
	var allowanceResponse model.AllowanceResponse
	_, err = c.do(req, &allowanceResponse)
	return &allowanceResponse, err
}

func (c *RestClient) SearchAreasByText(text string) (*model.AreasResponse, error) {
	req, err := c.newRequestWithParams(http.MethodGet, "./areas_search", nil, map[string]string{"text": text})
	if err != nil {
		return nil, err
	}
	var areasResponse model.AreasResponse
	_, err = c.do(req, &areasResponse)
	return &areasResponse, err
}

func (c *RestClient) newRequest(method, path string, body interface{}) (*http.Request, error) {
	return c.newRequestWithParams(method, path, body, map[string]string{})
}


func (c *RestClient) newRequestWithParams(method, path string, body interface{}, params map[string]string) (*http.Request, error) {
	rel := &url.URL{Path: path}
	u := c.baseUrl.ResolveReference(rel)
	if len(params) > 0 {
		q := u.Query()
		for param,value := range params {
			q.Set(param, value)
		}
		u.RawQuery = q.Encode()
	}
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Token", c.token)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.userAgent)
	return req, nil
}

func (c *RestClient) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(v)
	return resp, err
}

