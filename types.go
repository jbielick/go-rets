package main

// https://github.com/reso-workgroups/reso-synd

import (
	"encoding/xml"
	"gopkg.in/guregu/null.v3"
	"time"
)

// 2016-04-12T14:56:20+00:00
type Timestamp struct {
	time.Time
}

func (c *Timestamp) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	d.DecodeElement(&v, &start)
	parsed, err := time.Parse(time.RFC3339, v)
	if err != nil {
		return err
	}
	*c = Timestamp{parsed}
	return nil
}

// 2013-10-16T00:00:00Z
type xsDateTime struct {
	time.Time
}

func (c *xsDateTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	d.DecodeElement(&v, &start)
	parsed, err := time.Parse("2006-01-02T15:04:05Z", v)
	if err != nil {
		return err
	}
	*c = xsDateTime{parsed}
	return nil
}

type xsDate struct {
	time.Time
}

func (c *xsDate) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	d.DecodeElement(&v, &start)
	parsed, err := time.Parse("2006-01-02", v)
	if err != nil {
		return err
	}
	*c = xsDate{parsed}
	return nil
}

type Listings struct {
	Version          float32    `xml:"version,attr"`
	VersionTimestamp xsDateTime `xml:"versionTimestamp,attr"`
	ListingsKey      string     `xml:"listingsKey,attr"`
}

type SecureMoney struct {
	Value        string `xml:",chardata" json:",omitempty"`
	CurrencyCode string `xml:"currencyCode,attr" json:",omitempty"`
}

type FeeWithOptionalFrequency struct {
	SecureMoney
	CurrencyPeriod string `xml:"currencyPeriod,attr" json:",omitempty"`
}

type PriceWithOptionalFrequency struct {
	SecureMoney
	CurrencyPeriod string `xml:"currencyPeriod,attr" json:",omitempty"`
}

// commons:simpleTime is a mess
type OpenHouse struct {
	Date                  xsDate
	StartTime             string `json:",omitempty"`
	EndTime               string `json:",omitempty"`
	Description           string `json:",omitempty"`
	AppointmentRequiredYN null.Bool
	// OpenHouseAdditionalInformation commons:StringValueDescriptionPair
}

type Address struct {
	PreferenceOrder        null.Int `xml:"preference-order"`
	AddressPreferenceOrder null.Int `xml:"address-preference-order"`
	category               string   `json:",omitempty"`
	FullStreetAddress      string   `json:",omitempty"`
	StreetDirPrefix        string   `json:",omitempty"`
	StreetName             string   `json:",omitempty"`
	StreetSuffix           string   `json:",omitempty"`
	StreetDirSuffix        string   `json:",omitempty"`
	StreetAdditionalInfo   string   `json:",omitempty"`
	BoxNumber              string   `json:",omitempty"`
	UnitNumber             string   `json:",omitempty"`
	City                   string   `json:",omitempty"`
	StateOrProvince        string   `json:",omitempty"`
	PostalCode             string   `json:",omitempty"`
	CarrierRoute           string   `json:",omitempty"`
	Country                string   `json:",omitempty"`
	PrivacyType            string   `xml:"privacyType,attr" json:",omitempty"`
}

type Business struct {
	Name                          string
	BusinessId                    string   `json:",omitempty"`
	Phone                         string   `json:",omitempty"`
	Fax                           string   `json:",omitempty"`
	Email                         string   `json:",omitempty"`
	BusinessLeadRoutingEmail      string   `json:",omitempty"`
	WebsiteURL                    string   `json:",omitempty"`
	LogoURL                       string   `json:",omitempty"`
	Address                       *Address `json:",omitempty"`
	BusinessAdditionalInformation string   `json:",omitempty"`
}

type Area struct {
	Value             float32 `xml:",chardata"`
	AreaUnits         string  `xml:"areaUnits,attr" json:",omitempty"`
	MeasurementSource string  `xml:"measurementSource,attr" json:",omitempty"`
}

type OfficeCode struct {
	OfficeCodeId          string
	OfficeCodeDescription string `json:",omitempty"`
}

type Office struct {
	OfficeKey     string      `json:",omitempty"`
	OfficeId      string      `json:",omitempty"`
	Level         string      `json:",omitempty"`
	OfficeCode    *OfficeCode `json:",omitempty"`
	Name          string      `json:",omitempty"`
	CorporateName string      `json:",omitempty"`
	BrokerId      string      `json:",omitempty"`
	MainOfficeId  string      `json:",omitempty"`
	PhoneNumber   string      `json:",omitempty"`
	Fax           string      `json:",omitempty"`
	Address       *Address    `json:",omitempty"`
	OfficeEmail   string      `json:",omitempty"`
	Website       string      `json:",omitempty"`
	OfficeLogoURL string      `json:",omitempty"`
	// OfficeAdditionalInformation pair?
}

type AlternatePrice struct {
	AlternateListPrice    *PriceWithOptionalFrequency
	AlternateListPriceLow *PriceWithOptionalFrequency
}

type SchoolDistrict struct {
	DistrictName string `json:",omitempty"`
	// Elementary null.String
	// Middle null.String
	// JuniorHigh null.String
	// High null.String
	// DistrictURL null.String
	// DistrictPhoneNumber null.String
}

type Money struct {
	CurrencyCode string     `xml:"currencyCode,attr" json:",omitempty"`
	Value        null.Float `xml:",chardata"`
}

type License struct {
	LicenseExpDate  xsDateTime `json:",omitempty"`
	LicenseAmount   null.Float `json:",omitempty"`
	LicenseNumber   string     `json:",omitempty"`
	LicenseCategory string     `json:",omitempty"`
}

type ParticipantCode struct {
	ParticipantCodeId          string
	ParticipantCodeDescription string `json:",omitempty"`
}

type Participant struct {
	ParticipantKey      string
	ParticipantId       string           `json:",omitempty"`
	ParticipantCode     *ParticipantCode `json:",omitempty"`
	FirstName           string           `json:",omitempty"`
	LastName            string           `json:",omitempty"`
	Role                string           `json:",omitempty"`
	PrimaryContactPhone string           `json:",omitempty"`
	OfficePhone         string           `json:",omitempty"`
	MobilePhone         string           `json:",omitempty"`
	Email               string           `json:",omitempty"`
	Fax                 string           `json:",omitempty"`
	WebsiteURL          string           `json:",omitempty"`
	PhotoURL            string           `json:",omitempty"`
	Address             *Address         `json:",omitempty"`
	Licenses            []*License       `xml:"Licenses>License"`
}

type ValueWithOtherDescription struct {
	Value            string `xml:",chardata"`
	OtherDescription string `xml:"otherDescription,attr" json:",omitempty"`
}

type ArchitectureStyle struct {
	ValueWithOtherDescription
}

type PropertyType struct {
	ValueWithOtherDescription
}

type Media struct {
	MediaModificationTimestamp Timestamp
	MediaURL                   string
	MediaCaption               string `json:",omitempty"`
	MediaDescription           string `json:",omitempty"`
	MediaOrderNumber           null.Int
	MediaClassification        string `json:",omitempty"`
}

type Tax struct {
	Year           string `json:",omitempty"`
	Amount         *Money `json:",omitempty"`
	TaxDescription string `json:",omitempty"`
}

type Expense struct {
	ExpenseCategory string `json:",omitempty"`
	ExpenseValue    FeeWithOptionalFrequency
}

type MarketingInformation struct {
	PermitInternet               null.Bool
	PermitAddressOnInternet      null.Bool
	PermitPictureOnInternet      null.Bool
	PermitSignOnProperty         null.Bool
	HasSignOnProperty            null.Bool
	VOWEntireListingDisplay      null.Bool
	VOWAddressDisplay            null.Bool
	VOWAutomatedValuationDisplay null.Bool
	VOWConsumerComment           null.Bool
}

type DetailedCharacteristics struct {
	Appliances            []string `xml:"Appliances>Appliance" json:",omitempty"`
	ArchitectureStyle     *ArchitectureStyle
	HasAttic              null.Bool
	HasBarbecueArea       null.Bool
	HasBasement           null.Bool
	BuildingUnitCount     null.Int
	IsCableReady          null.Bool
	HasCeilingFan         null.Bool
	CondoFloorNum         null.Int
	CoolingSystems        []string `xml:"CoolingSystems>CoolingSystem"`
	HasDeck               null.Bool
	HasDisabledAccess     null.Bool
	HasDock               null.Bool
	HasDoorman            null.Bool
	HasDoublePaneWindows  null.Bool
	HasElevator           null.Bool
	ExteriorTypes         []string `xml:"ExteriorTypes>ExteriorType"`
	HasFireplace          null.Bool
	FloorCoverings        []string `xml:"FloorCoverings>FloorCovering"`
	HasGarden             null.Bool
	HasGatedEntry         null.Bool
	HasGreenhouse         null.Bool
	HeatingFuels          []string `xml:"HeatingFuels>HeatingFuel"`
	HeatingSystems        []string `xml:"HeatingSystems>HeatingSystem"`
	HasHotTubSpa          null.Bool
	Intercom              null.Bool
	HasJettedBathTub      null.Bool
	HasLawn               null.Bool
	LegalDescription      string
	HasMotherInLaw        null.Bool
	IsNewConstruction     null.Bool
	NumFloors             null.Float
	NumParkingSpaces      null.Int
	ParkingTypes          []string `xml:"ParkingTypes>ParkingType"`
	HasPatio              null.Bool
	HasPond               null.Bool
	HasPool               null.Bool
	HasPorch              null.Bool
	RoofTypes             []string `xml:"RoofTypes>RoofType"`
	RoomCount             null.Int
	Rooms                 []string `xml:"Rooms>Room"`
	HasRVParking          null.Bool
	HasSauna              null.Bool
	HasSecuritySystem     null.Bool
	HasSkylight           null.Bool
	HasSportsCourt        null.Bool
	HasSprinklerSystem    null.Bool
	HasVaultedCeiling     null.Bool
	ViewTypes             []string `xml:"ViewTypes>ViewType"`
	IsWaterfront          null.Bool
	HasWetBar             null.Bool
	WhatOwnerLoves        null.String
	IsWired               null.Bool
	YearUpdated           string `json:",omitempty"`
	AdditionalInformation string `json:",omitempty"`
}

type School struct {
	Name           string `json:",omitempty"`
	SchoolCategory string `json:",omitempty"`
	District       string `json:",omitempty"`
	Description    string `json:",omitempty"`
}

type Community struct {
	Subdivision string    `json:",omitempty"`
	Schools     []*School `xml:"Schools>School" json:",omitempty"`
}

type Neighborhood struct {
	Name        string `json:",omitempty"`
	Description string `json:",omitempty"`
}

type Location struct {
	Latitude           string     `json:",omitempty"`
	Longitude          string     `json:",omitempty"`
	Elevation          string     `json:",omitempty"`
	MapCoordinate      string     `json:",omitempty"`
	Directions         string     `json:",omitempty"`
	GeocodeOptions     string     `json:",omitempty"`
	County             string     `json:",omitempty"`
	StreetIntersection string     `json:",omitempty"`
	ParcelId           string     `json:",omitempty"`
	Community          *Community `json:",omitempty"`
	CommunityAmenities string     `json:",omitempty"`
	CommunityAddress   *Address   `json:",omitempty"`
	TotalNumFloors     null.Int
	Zoning             string `json:",omitempty"`
	BuildingAmenities  string `json:",omitempty"`
	BuildingUnitCount  null.Int
	Neighborhoods      []*Neighborhood `xml:"Neighborhoods>Neighborhood" json:",omitempty"`
}

type Listing struct {
	Id                      string                      `xml:"ListingKey" json:"id"`
	Address                 *Address                    `json:",omitempty"`
	ListPrice               *PriceWithOptionalFrequency `json:",omitempty"`
	ListPriceLow            *PriceWithOptionalFrequency `json:",omitempty"`
	AlternatePrices         []*AlternatePrice           `xml:"AlternatePrices>AlternatePrice" json:",omitempty"`
	ListingURL              string
	ProviderName            string
	ProviderURL             string
	ProviderCategory        string
	LeadRoutingEmail        string `json:",omitempty"`
	Bedrooms                null.Int
	Bathrooms               null.Int
	PropertyType            *PropertyType `json:",omitempty"`
	PropertySubType         *PropertyType `json:",omitempty"`
	ListingCategory         string        `json:",omitempty"`
	ListingStatus           string        `json:",omitempty"`
	MlsStatus               string        `json:",omitempty"`
	MarketingInformation    MarketingInformation
	Photos                  []*Media  `xml:"Photos>Photo" json:",omitempty"`
	DiscloseAddress         null.Bool `json:",omitempty"`
	ShortSale               null.Bool `json:",omitempty"`
	ListingDescription      string    `json:",omitempty"`
	MlsId                   string    `json:",omitempty"`
	MlsName                 string    `json:",omitempty"`
	MlsNumber               string    `json:",omitempty"`
	LivingArea              *Area     `json:",omitempty"`
	LotSize                 *Area     `json:",omitempty"`
	YearBuilt               string    `json:",omitempty"`
	ListingDate             xsDate
	TrackingItem            string `json:",omitempty"`
	ListingTitle            string `json:",omitempty"`
	FullBathrooms           null.Int
	ThreeQuarterBathrooms   null.Int
	HalfBathrooms           null.Int
	OneQuarterBathrooms     null.Int
	PartialBathrooms        null.Int
	ForeclosureStatus       string         `json:",omitempty"`
	ListingParticipants     []*Participant `xml:"ListingParticipants>Participant" json:",omitempty"`
	VirtualTours            []*Media       `xml:"VirtualTours>VirtualTour" json:",omitempty"`
	Videos                  []*Media       `xml:"Videos>Video" json:",omitempty"`
	Offices                 []*Office      `xml:"Offices>Office" json:",omitempty"`
	Brokerage               *Business      `json:",omitempty"`
	Location                *Location      `json:",omitempty"`
	Franchise               *Business      `json:",omitempty"`
	Builder                 *Business      `json:",omitempty"`
	PropertyManager         *Business      `json:",omitempty"`
	OpenHouses              []*OpenHouse   `xml:"OpenHouses>OpenHouse" json:",omitempty"`
	Taxes                   []*Tax         `xml:"Taxes>Tax" json:",omitempty"`
	Expenses                []*Expense     `xml:"Expenses>Expense" json:",omitempty"`
	DetailedCharacteristics DetailedCharacteristics
	ModificationTimestamp   Timestamp
	Disclaimer              string
}
