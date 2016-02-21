package goeasypost

import (
)

func NewTracker(trackingNumber string, carrier string) (*Tracker, error) {

	var params = make(map[string]string)

	params["tracker[tracking_code]"] = trackingNumber
	params["tracker[carrier]"] = carrier

	var tracker = new(Tracker)

	err := getResource("POST", "/trackers", params, tracker)
	if err != nil {
		return nil, err
	}



	return tracker, nil
}

func GetTrackerList() (*TrackerList, error) {

	var params = make(map[string]string)

	var tracker = new(TrackerList)

	err := getResource("GET", "/trackers", params, tracker)
	if err != nil {
		return nil, err
	}



	return tracker, nil
}
