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

func NewTracker(trackingNumber string, carrier string) (*Tracker, error) {

	var params = make(map[string]string)

	params["tracker[tracking_code]"] = trackingNumber
	params["tracker[carrier]"] = carrier

	var tracker = new(Tracker)

	err := createEntity("trackers", params, tracker)
	if err != nil {
		return nil, err
	}

	return tracker, nil
}

func GetTrackerList(filterParams map[string]string) (*TrackerList,  error) {


	var tracker = new(TrackerList)

	err := getEntityCollection("trackers", filterParams, tracker)
	if err != nil {
		return nil, err
	}

	return tracker, nil
}

func GetTracker(id string) (*Tracker, error) {

	tracker := &Tracker{}

	err := getEntityById("trackers", id, tracker)
	if err != nil {
		return nil, err
	}

	return tracker, nil
}
