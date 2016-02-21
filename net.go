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

func createEntity( rootPath string, params map[string]string, respStruct interface{}) ( error) {

	err := getResource("POST", fmt.Sprintf("/%s", rootPath), params, respStruct)
	if err != nil {
		return err
	}

	return nil

}


func getEntityById( rootPath string, id string, respStruct interface{}) (  error) {

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
		epet := new(EasyPostErrorTest)
		err = json.Unmarshal(data, &epet)
		if err != nil {
			return errors.New(fmt.Sprintf("Unmarshal failure %s", err))
		}

		if epet.ErrorResponse != nil {
			return  errors.New(fmt.Sprintf("Easypost Error Code: %s Message: %s Errors: %v", epet.ErrorResponse.Code, epet.ErrorResponse.Message, epet.ErrorResponse.Errors))
		}

		return errors.New(fmt.Sprintf("Invalid http status code %d received from %s%s", resp.StatusCode, HostUri, uriSuffix))
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New(fmt.Sprintf("Read response data failure %s", err))
	}
	/*var prettyJSON bytes.Buffer
	json.Indent(&prettyJSON, data, "", "\t")


	logme.Println("Raw response:", string(prettyJSON.Bytes()))
*/
	epet := new(EasyPostErrorTest)
	err = json.Unmarshal(data, &epet)
	if err != nil {
		return errors.New(fmt.Sprintf("Unmarshal failure %s", err))
	}

	if epet.ErrorResponse != nil {
		return  errors.New(fmt.Sprintf("Easypost Error Code: %s Message: %s Errors: %v", epet.ErrorResponse.Code, epet.ErrorResponse.Message, epet.ErrorResponse.Errors))
	}

	err = json.Unmarshal(data, respStruct)
	if err != nil {
		return errors.New(fmt.Sprintf("Unmarshal failure %s", err))
	}

	return nil
}
