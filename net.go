package goeasypost
import (
	"net/http"
	"fmt"
	"io/ioutil"
	"errors"
	"net/url"
	"bytes"
	"encoding/json"
)




func GetResource(method, uriSuffix string, params map[string]string, respStruct interface{}) (error) {

	paramData := url.Values{}
	for k, v := range params {
		logme.Println(fmt.Sprintf("PARAM->%v:%v",k,v))
		paramData.Add(k,v)
	}

	urlStr := fmt.Sprintf("%v%v", HostUri,uriSuffix)


	logme.Println(fmt.Sprintf("URL: %v",urlStr))
	req, err := http.NewRequest(method, urlStr,  bytes.NewBufferString(paramData.Encode()))

	req.SetBasicAuth(ApiKey, "")

	req.URL.Query()

	cli := &http.Client{}

	resp, err := cli.Do(req)
	if (resp != nil) {
		defer resp.Body.Close()
	}
	if (err != nil) {
		return err
	}
	logme.Printf("HTTP RESPONSE CODE: %d", resp.StatusCode)

	if(resp.StatusCode<200 || resp.StatusCode>300)  {

		for k, v := range resp.Header {
			logme.Println("name:", k, "value:", v)
		}

		return errors.New(fmt.Sprintf("Invalid http status code %d received from %s%s",resp.StatusCode, HostUri, uriSuffix))
	}

	data, err := ioutil.ReadAll(resp.Body)


	err = json.Unmarshal(data, &respStruct)
	if(err!=nil) {
		return err
	}

	return nil
}

