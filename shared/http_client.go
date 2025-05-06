package shared

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type HttpClient struct {
	method              string
	url                 string
	authorizationHeader string
	queryParamCount     int
	form                url.Values
	response            *http.Response
}

func NewHttpClientGET(url string) *HttpClient {
	return &HttpClient{
		method:              "GET",
		url:                 url,
		authorizationHeader: "",
		queryParamCount:     0,
	}
}

func NewHttpClientPOST(url string) *HttpClient {
	return &HttpClient{
		method:              "POST",
		url:                 url,
		authorizationHeader: "",
		queryParamCount:     0,
	}
}

func NewHttpClientPATCH(url string) *HttpClient {
	return &HttpClient{
		method:              "PATCH",
		url:                 url,
		authorizationHeader: "",
		queryParamCount:     0,
	}
}

func NewHttpClientDELETE(url string) *HttpClient {
	return &HttpClient{
		method:              "DELETE",
		url:                 url,
		authorizationHeader: "",
		queryParamCount:     0,
	}
}

func (client *HttpClient) GetResponse() *http.Response {
	return client.response
}

func (client *HttpClient) SetAuthorizationHeaderBasicAuth(username string, password string) {
	auth := username + ":" + password
	encodedAuth := base64.StdEncoding.EncodeToString([]byte(auth))
	client.authorizationHeader = "Basic " + encodedAuth
}

func (client *HttpClient) SetAuthorizationHeaderBearerToken(token string) {
	client.authorizationHeader = "Bearer " + token
}

func (client *HttpClient) AddQueryParam(key string, value any) {
	valueString := fmt.Sprintf("%v", value)
	if client.queryParamCount <= 0 {
		client.url += "?" + key + "=" + valueString
		client.queryParamCount += 1
		return
	}
	client.url += "&" + key + "=" + valueString
	client.queryParamCount += 1
}

func (client *HttpClient) AddFormValue(key string, value any) {
	if client.form == nil {
		client.form = url.Values{}
	}
	valueString := fmt.Sprintf("%v", value)
	client.form.Add(key, valueString)
}

func (client *HttpClient) SendAndParseBody(responseBody any) error {
	bodyReader := client.getBodyReader()
	req, err := http.NewRequest(client.method, client.url, bodyReader)
	if err != nil {
		return err
	}

	if bodyReader != nil {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	if client.authorizationHeader != "" {
		req.Header.Add("Authorization", client.authorizationHeader)
	}

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	client.response = resp

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &responseBody)
	if err != nil {
		return err
	}

	return nil
}

func (client *HttpClient) getBodyReader() io.Reader {
	if client.method == "GET" {
		return nil
	}

	return strings.NewReader(client.form.Encode())
}
