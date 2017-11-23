// Copyright 2017 Marin Basic <marin@marin-basic.com>. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package gomercedesdealer

// dealer with navigation links.
type HalifiedDealers struct {
	Links   DefaultLinks `json:"_links"`
	Dealers []Dealer     `json:"dealers"`
}

// default links
type DefaultLinks struct {
	Self     Link `json:"self"`
	Next     Link `json:"next"`
	Previous Link `json:"previous"`
}

// the representation of a link
type Link struct {
	Href string `json:"href"`
}

// self links
type SelfLink struct {
	Self Link `json:"self"`
}

// Code: brand qualifier
// Name: brand name
type Brand struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

// Street: street and house number, if applicable
// AddressAddition: address addition
// ZipCode: postal code
// City: city
// District: district, area of city
// CountryIsoCode: ISO A2 country code
type Address struct {
	Street          string `json:"street"`
	AddressAddition string `json:"addressAddition"`
	ZipCode         string `json:"zipCode"`
	City            string `json:"city"`
	District        string `json:"district"`
	CountryIsoCode  string `json:"countryIsoCode"`
}

// RegionRepresentation
type RegionRepresentation struct {
	Region    string `json:"region"`
	SubRegion string `json:"subRegion"`
}

// GeoCoordinates in decimal degrees
type GeoCoordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// CommunicationChannels
type CommunicationChannels struct {
	Fax        string `json:"fax"`
	Email      string `json:"email"`
	Website    string `json:"website"`
	Facebook   string `json:"facebook"`
	Mobile     string `json:"mobile"`
	GooglePlus string `json:"googlePlus"`
	Twitter    string `json:"twitter"`
	Phone      string `json:"phone"`
}

// Value: numeric amount of distance
// Unit: unit of length [M, KM, MILE]
type TheDistance struct {
	Value float64 `json:"value"`
	Unit  string  `json:"unit"`
}

// Function
type Function struct {
	Brand        Brand                   `json:"brand"`
	ProductGroup ProductGroupOfAFunction `json:"productGroup"`
	Activity     ActivityOfAFunction     `json:"activity"`
}

// Code: The product group code, valid values are: PASSENGER-CAR, VAN
// Name: product group name
type ProductGroupOfAFunction struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

// Code: The activity of the dealer, valid actitvity values are
//    PARTS -> Spare Parts Sales
//    SALES -> Sales of new vehicles
//    SERVICE -> Maintaining and repair
//    USED-VEHICLES-TRADE -> Sales of used vehicles
// Name: the corresponding activity name
type ActivityOfAFunction struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

// FunctionOpeningHours
type FunctionOpeningHours struct {
	Function Function `json:"function"`
	Weekdays Weekdays `json:"weekdays"`
}

// opening hours for a week, represented by a map where the key
// is a weekday and the value is an object containing the time periods. There
// are a maximum of two time slots for each day. If there exists no opening
// hours for one day, they will be shown as closed.
type Weekdays struct {
	Monday    WeekdayInfo `json:"MONDAY"`
	Tuesday   WeekdayInfo `json:"TUESDAY"`
	Wednesday WeekdayInfo `json:"WEDNESDAY"`
	Thursday  WeekdayInfo `json:"THURSDAY"`
	Friday    WeekdayInfo `json:"FRIDAY"`
	Saturday  WeekdayInfo `json:"SATURDAY"`
	Sunday    WeekdayInfo `json:"SUNDAY"`
}

// Status: OPEN or CLOSED
type WeekdayInfo struct {
	Status  string   `json:"status"`
	Periods []Period `json:"periods"`
}

// Period
type Period struct {
	From  string `json:"from"`
	Until string `json:"until"`
}

// the representation of an dealer
type Dealer struct {
	Links          SelfLink               `json:"_links"`
	DealerId       string                 `json:"dealerId"`
	LegalName      string                 `json:"legalName"`
	NameAddition   string                 `json:"nameAddition"`
	Brands         []Brand                `json:"brands"`
	Address        Address                `json:"address"`
	Region         RegionRepresentation   `json:"region"`
	GeoCoordinates GeoCoordinates         `json:"geoCoordinates"`
	Communication  CommunicationChannels  `json:"communicationChannels"`
	Distance       TheDistance            `json:"distance"`
	Functions      []Function             `json:"functions"`
	OpeningHours   []FunctionOpeningHours `json:"openingHours"`
}

// countries and navigation links for additional results.
type HalifiedCountries struct {
	Links     DefaultLinks `json:"_links"`
	Countries []Country    `json:"countries"`
}

// Country
type Country struct {
	CountryId string `json:"countryId"`
	IsoCode   string `json:"isoCode"`
	Name      string `json:"name"`
}

// DealerParam used for filtering request
type DealerParam struct {
	DealerIds      []string
	Latitude       float64
	Longitude      float64
	RadiusValue    int
	RadiusUnit     string
	CountryIsoCode string
	City           string
	Name           string
	Brand          string
	ProductGroup   string
	Activity       string
	Fields         string
	Page           int
	PageSize       int
}
