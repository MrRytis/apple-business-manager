package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Client interface {
	BulkEnroll(req *BulkEnrollRequest) (*BulkEnrollResponse, error)
	CheckStatus(req *StatusRequest) (*StatusResponse, error)
}

type AbmClient struct {
	url    string
	secret string
	client *http.Client
}

func NewAbmClient(url, secret string) *AbmClient {
	return &AbmClient{
		url:    url,
		secret: secret,
		client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

func (a *AbmClient) BulkEnroll(r *BulkEnrollRequest) (*BulkEnrollResponse, error) {
	jsonData, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("%s/bulk-enroll-devices", a.url),
		bytes.NewBuffer(jsonData),
	)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := a.client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	var res BulkEnrollResponse
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (a *AbmClient) CheckStatus(r *StatusRequest) (*StatusResponse, error) {
	jsonData, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("%s/check-transaction-status", a.url),
		bytes.NewBuffer(jsonData),
	)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := a.client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	var res StatusResponse
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
