package goeasypost

import (
)

func RetrieveAddress(id string) (*Address, error) {


	address := &Address{}

	err := getEntityById("addresses", id, address)
	if err != nil {
		return nil, err
	}

	return address, nil
}

func NewAddress(reqAddr *Address) (*Address, error) {

	params := make(map[string]string)

	err := flattenStructMap(reqAddr, "address", params)
	if err != nil {
		return nil, err
	}

	if reqAddr.VerifyRequest != nil {
		if  reqAddr.VerifyRequest.Strict {
			params["verify_strict[]"] = reqAddr.VerifyRequest.Type
		} else {
			params["verify[]"] = reqAddr.VerifyRequest.Type
		}
	}

	var address = new(Address)

	err = createEntity("addresses", params, address)
	if err != nil {
		return nil, err
	}

	return address, nil
}
