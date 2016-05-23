package main

type PriceWithOptionalFrequency struct {
	Value          string `xml:",chardata"`
	CurrencyPeriod string `xml:"currencyPeriod,attr"`
	CurrencyCode   string `xml:"currencyCode,attr"`
}

type Address struct {
	PreferenceOrder        int `xml:"preference-order"`
	AddressPreferenceOrder int `xml:"address-preference-order"`
	FullStreetAddress      string
	StreetDirPrefix        string
	StreetName             string
	StreetSuffix           string
	StreetDirSuffix        string
	StreetAdditionalInfo   string
	BoxNumber              string
	UnitNumber             string
	City                   string
	StateOrProvince        string
	PostalCode             string
	Country                string
	CarrierRoute           string
	PrivacyType            string `xml:privacyType,attr"`
}

type Business struct {
	Name                          string
	BusinessId                    string
	Phone                         string
	Fax                           string
	Email                         string
	WebsiteURL                    string
	LogoURL                       string
	Address                       Address
	BusinessAdditionalInformation string
}

type Area struct {
	Value             float32 `xml:",chardata"`
	AreaUnits         string  `xml:"areaUnits,attr"`
	MeasurementSource string  `xml:"measurementSource,attr"`
}

type Office struct {
	OfficeKey string
	OfficeId  string
	Level     string
	// OfficeCode OfficeCode
	Name          string
	CorporateName string
	BrokerId      string
	MainOfficeId  string
	PhoneNumber   string
	Fax           string
	Address       Address
	OfficeEmail   string
	Website       string
	OfficeLogoURL string
	// OfficeAdditionalInformation pair?
}

type Participant struct {
	ParticipantKey string
	ParticipantId  string
	// ParticipantCode ParticipantCode
	FirstName           string
	LastName            string
	Role                string
	PrimaryContactPhone string
	OfficePhone         string
	MobilePhone         string
	Email               string
	Fax                 string
	WebsiteURL          string
	PhotoURL            string
	Address             Address
	// Licenses []License `xml:"Licenses>License"`
}

type ArchitectureStyle struct {
	Value            string `xml:",chardata"`
	OtherDescription string `xml:"otherDescription,attr"`
}

type Media struct {
	MediaModificationTimestamp customTime `xml:",omitempty"`
	MediaURL                   string
}

type DetailedCharacteristics struct {
	ArchitectureStyle     string
	HasAttic              bool
	Appliances            []string `xml:"Appliances>Appliance"`
	HasBarbecueArea       bool
	HasBasement           bool
	BuildingUnitCount     int
	IsCableReady          bool
	HasCeilingFan         bool
	CondoFloorNum         int
	CoolingSystems        []string `xml:"CoolingSystems>CoolingSystem"`
	HasDeck               bool
	HasDisabledAccess     bool
	HasDock               bool
	HasDoorman            bool
	HasDoublePaneWindows  bool
	HasElevator           bool
	ExteriorTypes         []string `xml:"ExteriorTypes>ExteriorType"`
	HasFireplace          bool
	FloorCoverings        []string `xml:"FloorCoverings>FloorCovering"`
	HasGarden             bool
	HasGatedEntry         bool
	HasGreenhouse         bool
	HeatingFuels          []string `xml:"HeatingFuels>HeatingFuel"`
	HeatingSystems        []string `xml:"HeatingSystems>HeatingSystem"`
	HasHotTubSpa          bool
	Intercom              bool
	HasJettedBathTub      bool
	HasLawn               bool
	LegalDescription      string
	HasMotherInLaw        bool
	IsNewConstruction     bool
	NumFloors             float32  `xml:",omitempty"`
	NumParkingSpaces      int      `xml:",omitempty"`
	ParkingTypes          []string `xml:"ParkingTypes>ParkingType"`
	HasPatio              bool
	HasPond               bool
	HasPool               bool
	HasPorch              bool
	RoofTypes             []string `xml:"RoofTypes>RoofType"`
	RoomCount             int
	Rooms                 []string `xml:"Rooms>Room"`
	HasRVParking          bool
	HasSauna              bool
	HasSecuritySystem     bool
	HasSkylight           bool
	HasSportsCourt        bool
	HasSprinklerSystem    bool
	HasVaultedCeiling     bool
	ViewTypes             []string `xml:"ViewTypes>ViewType"`
	IsWaterfront          bool
	HasWetBar             bool
	WhatOwnerLoves        string
	IsWired               bool
	YearUpdated           string `xml:",omitempty"`
	AdditionalInformation string
}

type Neighborhood struct {
	Name        string
	Description string
}

type Location struct {
	Latitude           string
	Longitude          string
	Elevation          string
	MapCoordinate      string
	Directions         string
	GeocodeOptions     string
	County             string
	StreetIntersection string
	ParcelId           string
	Community          string
	CommunityAmenities string
	// CommunityAddress Address
	TotalNumFloors    int
	Zoning            string
	BuildingAmenities string
	BuildingUnitCount int
	Neighborhoods     []Neighborhood `xml:"Neighborhoods>Neighborhood"`
}

type Listing struct {
	Id               string `xml:"ListingKey" json:"id"`
	Address          Address
	ListPrice        PriceWithOptionalFrequency
	ListingURL       string
	ProviderName     string
	ProviderURL      string
	ProviderCategory string
	LeadRoutingEmail string
	Bedrooms         int
	Bathrooms        int
	PropertyType     string
	PropertySubType  string
	ListingCategory  string
	ListingStatus    string
	// MarketingInformation
	Photos             []Media `xml:"Photos>Photo"`
	DiscloseAddress    bool
	ShortSale          bool
	ListingDescription string
	MlsId              string
	MlsName            string
	MlsNumber          string
	LivingArea         Area
	// LotSizeAreaUnits      string  `xml:"areaUnits,attr"`
	LotSize               Area
	YearBuilt             string
	ListingDate           string
	TrackingItem          string
	ListingTitle          string
	FullBathrooms         int
	ThreeQuarterBathrooms int
	HalfBathrooms         int
	OneQuarterBathrooms   int
	PartialBathrooms      int
	ForeclosureStatus     string
	ListingParticipants   []Participant `xml:"ListingParticipants>ListingParticipant"`
	// VirtualTours []Media
	Videos          []Media  `xml:"Videos>Video"`
	Offices         []Office `xml:"Offices>Office"`
	Brokerage       Business `xml:",omitempty"`
	Franchise       Business `xml:",omitempty"`
	Builder         Business `xml:",omitempty"`
	PropertyManager Business
	// OpenHouses []OpenHouse
	// Taxes []Tax
	Expenses                []string
	DetailedCharacteristics DetailedCharacteristics
	ModificationTimestamp   customTime `xml:",omitempty"`
	Disclaimer              string
}
