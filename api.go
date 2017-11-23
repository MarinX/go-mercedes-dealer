// Copyright 2017 Marin Basic <marin@marin-basic.com>. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package gomercedesdealer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	API_URL = "https://api.mercedes-benz.com/dealer/v1"
)

type API struct {
	key    string
	client *http.Client
}

// Returns new dealer API
func New(key string) *API {
	return &API{
		key:    key,
		client: &http.Client{},
	}
}

// This request returns the countries supported by the Dealer API
func (a *API) GetCountries(page int, pageSize int) (*HalifiedCountries, error) {
	hc := new(HalifiedCountries)

	params := make(map[string]string)
	if page > 0 {
		params["page"] = fmt.Sprintf("%d", page)
	}

	if pageSize > 0 {
		params["pageSize"] = fmt.Sprintf("%d", pageSize)
	}

	if err := a.call("/countries", params, hc); err != nil {
		return nil, err
	}

	return hc, nil
}

// This request returns dealers (dealer, garage, retailer, etc.) for given parameters
func (a *API) GetDealers(dealerParam *DealerParam) (*HalifiedDealers, error) {
	hd := new(HalifiedDealers)
	params := make(map[string]string)
	if dealerParam != nil {
		params = mapFromStruct(dealerParam)
	}

	if err := a.call("/dealers", params, hd); err != nil {
		return nil, err
	}

	return hd, nil
}

// This request returns single dealer by id eg GS0038375
func (a *API) GetDealer(dealerId string) (*Dealer, error) {
	d := new(Dealer)
	if err := a.call(fmt.Sprintf("/dealers/%s", dealerId), nil, d); err != nil {
		return nil, err
	}
	return d, nil
}

// Private function to call mercedes-benz dealer API and parse response into model
func (a *API) call(path string, params map[string]string, responseModel interface{}) error {
	req, err := http.NewRequest("GET", API_URL+path, nil)
	if err != nil {
		return err
	}

	q := req.URL.Query()
	for key, val := range params {
		q.Add(key, val)
	}

	q.Add("apikey", a.key)
	req.URL.RawQuery = q.Encode()
	req.Header.Set("Content-Type", "application/json")

	resp, err := a.client.Do(req)
	if err != nil {
		return err
	}

	// visit https://developer.mercedes-benz.com/apis/dealer_api/docs under "General error handling"
	switch resp.StatusCode {
	case 200:
		// all OK, continue
		break
	case 404:
		return fmt.Errorf("404:The Dealer for the requested ID could not be found")
	case 204:
		return fmt.Errorf("204:No content could be found for the requested search parameters")
	case 400:
		return fmt.Errorf("400:An input parameter was not provided correctly")
	case 500:
		return fmt.Errorf("500:An error occurred on the server side")
	default:
		return fmt.Errorf("%d:Unknown HTTP error code", resp.StatusCode)
	}

	buff, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Parse JSON into provided model
	return json.Unmarshal(buff, responseModel)
}
