/*
 * Copyright (c) 2016 Stewart Buskirk <mitymauser@gmail.com>
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */
package goeasypost

import ()

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
		if reqAddr.VerifyRequest.Strict {
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
