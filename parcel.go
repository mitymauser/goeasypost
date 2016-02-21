package goeasypost

import (
)

func RetrieveParcel(id string) (*Parcel, error) {


	parcel := &Parcel{}


	err := getEntityById("parcels", id , parcel)
	if err != nil {
		return nil, err
	}

	return parcel, nil
}

func NewParcel(reqParcel *Parcel) (*Parcel, error) {

	params := make(map[string]string)

	err := flattenStructMap(reqParcel, "parcel", params)
	if err != nil {
		return nil, err
	}

	var parcel = new(Parcel)

	err = createEntity("parcels", params , parcel)
	if err != nil {
		return nil, err
	}

	return parcel, nil
}

