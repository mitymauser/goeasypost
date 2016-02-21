package goeasypost

import "time"

type EasyPostErrorTest struct {
	ErrorResponse *EasyPostError `json:"error"`
}


type EasyPostError struct {
	Code    string        `json:"code"`
	Errors  []interface{} `json:"errors"`
	Message string        `json:"message"`
}

type Fee struct {
	Type     string `json:"type"`
	Amount   string `json:"amount"`
	Charged  bool   `json:"charges"`
	Refunded bool   `json:"refunded"`
}

type TrackerList struct {
	Trackers      []Tracker      `json:"trackers"`
	HasMore       bool           `json:"has_more"`
}

type TrackingLocation struct {
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
	Zip     string `json:"zip"`
}

type TrackingDetail struct {
	Status           string           `json:"status"`
	Message          string           `json:"message"`
	Datetime         time.Time        `json:"datetime"`
	TrackingLocation TrackingLocation `json:"tracking_location"`
}

type CarrierDetail struct {
	Service              string     `json:"service"`
	ContainerType        string     `json:"container_type"`
	EstDeliveryDateLocal *time.Time `json:"est_delivery_date_local"`
	EstDeliveryTimeLocal *time.Time `json:"est_delivery_time_local"`
}

type Tracker struct {
	Id              string           `json:"id"`
	Mode            string           `json:"mode"`
	TrackingCode    string           `json:"tracking_code"`
	Status          string           `json:"status"`
	ShipmentId      string           `json:"shipment_id"`
	Carrier         string           `json:"carrier"`
	TrackingDetails []TrackingDetail `json:"tracking_details"`
	Weight          float32          `json:"weight"`
	EstDeliveryDate *time.Time       `json:"est_delivery_date"`
	CreatedAt       time.Time        `json:"created_at"`
	UpdatedAt       time.Time        `json:"updated_at"`
	SignedBy        string           `json:"signed_by"`
	CarrierDetail   CarrierDetail    `json:"carrier_detail"`
	Fees            []Fee            `json:"fees"`
}

type Address struct {
	VerifyRequest *VerifyRequest     `json:"-"`
	Id                string         `json:"id"`
	Object            string         `json:"object"`
	Mode              string         `json:"mode"`
	Street1           string         `json:"street1"`
	Street2           string         `json:"street2"`
	City              string         `json:"city"`
	State             string         `json:"state"`
	Zip               string         `json:"zip"`
	Country           string         `json:"country"`
	Residential       bool           `json:"residential"`
	CarrierFacility   string         `json:"carrier_facility"`
	Name              string         `json:"name"`
	Company           string         `json:"company"`
	Phone             string         `json:"phone"`
	Email             string         `json:"email"`
	FederalTaxId      string         `json:"federal_tax_id"`
	StateTaxId        string         `json:"state_tax_id"`
	Verifications     *Verifications `json:"verifications"`
}

type VerifyRequest struct {
	Strict bool
	Type string
}

type Verifications struct {
	Zip4     *Verification `json:"zip4"`
	Delivery *Verification `json:"delivery"`
}

type Verification struct {
	Success bool     `json:"success"`
	Errors  [][][]string `json:"errors"`
}

type Parcel struct {
	Id                string         `json:"id"`
	Object            string         `json:"object"`
	Mode              string         `json:"mode"`
	Length 	float32        `json:"length"`
	Width 	float32         `json:"width"`
	Height 	float32          `json:"height"`
	PredefinedPackage 	string    `json:"predefined_package"`
	Weight 	float32                       `json:"weight"`
	CreatedAt 	time.Time                          `json:"created_at"`
	UpdatedAt 	time.Time                          `json:"updated_at"`
}
