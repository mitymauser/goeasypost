package goeasypost

import (
	"time"
	"fmt"
	"errors"
)

type TrackerList struct {
	ErrorResponse *EasyPostError `json:"error"`
	Trackers      []interface{}  `json:"trackers"`
	HasMore       bool           `json:"has_more"`
}

type TrackingLocation struct {
	City    string              `json:"city"`
	State   string              `json:"state"`
	Country string              `json:"country"`
	Zip     string              `json:"zip"`

}

type TrackingDetail struct {
	Status           string              `json:"status"`
	Message          string              `json:"message"`
	Datetime         time.Time           `json:"datetime"`
	TrackingLocation TrackingLocation    `json:"trackingLocation"`
}

type CarrierDetail struct {
	Service              string    `json:"service"`
	ContainerType        string    `json:"containerType"`
	EstDeliveryDateLocal string    `json:"estDeliveryDateLocal"`
	EstDeliveryTimeLocal string    `json:"estDeliveryTimeLocal"`
}

type Tracker struct {
	ErrorResponse *EasyPostError                   `json:"error"`
	Id              string                         `json:"id"`
	Mode            string                         `json:"mode"`
	TrackingCode    string                         `json:"trackingCode"`
	Status          string                         `json:"status"`
	ShipmentId      string                         `json:"shipmentId"`
	Carrier         string                         `json:"carrier"`
	TrackingDetails []interface{}                  `json:"trackingDetails"`
	Weight          float32                        `json:"weight"`
	EstDeliveryDate time.Time                      `json:"estDeliveryDate"`
	SignedBy        string                         `json:"signedBy"`
	CarrierDetail   CarrierDetail                  `json:"carrierDetail"`
}


func CreateTracker(trackingNumber string, carrier string) (*Tracker, error) {

	var params =  make(map[string]string)

	params["tracker[tracking_code]"]=trackingNumber
	params["tracker[carrier]"]=carrier

	var tracker = new(Tracker)

	err := GetResource("POST","/trackers",params,tracker)
	if(err != nil) {
		return nil, err
	}

	if(tracker.ErrorResponse != nil) {
		return nil, errors.New(fmt.Sprintf("Easypost Error Code: %s Message %s Errors: %v", tracker.ErrorResponse.Code,tracker.ErrorResponse.Message, tracker.ErrorResponse.Errors))
	}

	return tracker, nil
}