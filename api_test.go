// Copyright 2017 Marin Basic <marin@marin-basic.com>. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package gomercedesdealer

import (
	"strings"
	"testing"
)

func TestApiCalls(t *testing.T) {

	// create new api
	api := New("Yjzo77fTyM6mZTnHnQ6lAzk8Iu6R7gCq")

	// get all countries
	countries, err := api.GetCountries(1, 5)
	if err != nil {
		t.Error("[GetCountries]", err)
		return
	}

	// check the param pageSize if it works
	if len(countries.Countries) != 5 {
		t.Error("[GetCountries] Size of countries are not 5", len(countries.Countries))
		return
	}

	// Find all dealers near Zagreb
	dealers, err := api.GetDealers(&DealerParam{
		City: "Zagreb",
	})
	if err != nil {
		t.Error("[GetDealers]", err)
		return
	}

	found := false
	// there should be a dealer Emil Frey with ID of GS0038375
	// lets find it
	for _, val := range dealers.Dealers {
		if val.DealerId == "GS0038375" {
			found = true
		}
	}

	if !found {
		t.Error("[GetDealers] Did not found dealer in Zagreb ")
		return
	}

	// Get single dealer - code for Emil Frey dealer in Zagreb
	singleDealer, err := api.GetDealer("GS0038375")
	if err != nil {
		t.Error("[GetDealer]", err)
		return
	}

	if !strings.Contains(singleDealer.LegalName, "Emil") {
		t.Error("[GetDealer] Dealer did not match for provided param ", *singleDealer)
	}
}
