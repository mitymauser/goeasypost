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

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)


type easyPostErrorTest struct {
	errorResponse *easyPostError `json:"error"`
}

type easyPostError struct {
	code    string        `json:"code"`
	errors  []interface{} `json:"errors"`
	message string        `json:"message"`
}

func callEntityWithCommand(rootPath string, id string, command string, params map[string]string, respStruct interface{}) error {

	err := getResource("POST", fmt.Sprintf("/%s/%s/%s", rootPath, id, command), params, respStruct)
	if err != nil {
		return err
	}

	return nil

}

func createEntity(rootPath string, params map[string]string, respStruct interface{}) error {

	err := getResource("POST", fmt.Sprintf("/%s", rootPath), params, respStruct)
	if err != nil {
		return err
	}

	return nil

}

func getEntityCollection(rootPath string, params map[string]string, respStruct interface{}) error {

	err := getResource("GET", fmt.Sprintf("/%s", rootPath), params, respStruct)
	if err != nil {
		return err
	}

	return nil

}

func getEntityById(rootPath string, id string, respStruct interface{}) error {

	err := getResource("GET", fmt.Sprintf("/%s/%s", rootPath, id), nil, respStruct)
	if err != nil {
		return err
	}

	return nil

}

func getResource(method, uriSuffix string, params map[string]string, respStruct interface{}) error {

	paramData := url.Values{}
	for k, v := range params {
		logme.Println(fmt.Sprintf("PARAM->%v:%v", k, v))
		paramData.Add(k, v)
	}

	urlStr := fmt.Sprintf("%v%v", HostUri, uriSuffix)

	logme.Println(fmt.Sprintf("URL: %s %v", method, urlStr))
	req, err := http.NewRequest(method, urlStr, bytes.NewBufferString(paramData.Encode()))

	req.SetBasicAuth(ApiKey, "")

	cli := &http.Client{}

	resp, err := cli.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return err
	}
	logme.Printf("HTTP RESPONSE CODE: %d", resp.StatusCode)

	if resp.StatusCode < 200 || resp.StatusCode > 300 {

		for k, v := range resp.Header {
			logme.Println("name:", k, "value:", v)
		}

		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return errors.New(fmt.Sprintf("Read response data failure %s", err))
		}
		epet := new(easyPostErrorTest)
		err = json.Unmarshal(data, &epet)
		if err != nil {
			return errors.New(fmt.Sprintf("Unmarshal failure %s", err))
		}

		if epet.errorResponse != nil {
			return errors.New(fmt.Sprintf("Easypost Error Code: %s Message: %s Errors: %v", epet.errorResponse.code, epet.errorResponse.message, epet.errorResponse.errors))
		}

		return errors.New(fmt.Sprintf("Invalid http status code %d received from %s%s", resp.StatusCode, HostUri, uriSuffix))
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New(fmt.Sprintf("Read response data failure %s", err))
	}

		//when debugging, displaying the raw JSON returned is highly useful
		var prettyJSON bytes.Buffer
		json.Indent(&prettyJSON, data, "", "\t")


		logme.Println("Raw response:", string(prettyJSON.Bytes()))

	epet := new(easyPostErrorTest)
	err = json.Unmarshal(data, &epet)
	if err != nil {
		return errors.New(fmt.Sprintf("Unmarshal failure %s", err))
	}

	if epet.errorResponse != nil {
		return errors.New(fmt.Sprintf("Easypost Error Code: %s Message: %s Errors: %v", epet.errorResponse.code, epet.errorResponse.message, epet.errorResponse.errors))
	}

	err = json.Unmarshal(data, respStruct)
	if err != nil {
		return errors.New(fmt.Sprintf("Unmarshal failure %s", err))
	}

	return nil
}
