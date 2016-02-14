package goeasypost

import (
	"testing"
	"os"
	"flag"
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
	"encoding/json"
	"github.com/mitymauser/goeasypost/models"
)


func TestMain(m *testing.M) {
	flag.Parse()
	log.SetOutput(os.Stdout)
	os.Exit(m.Run())
}


func TestHello(t *testing.T) {




	req, err := http.NewRequest("GET", "https://api.easypost.com/v2/trackers", nil)
	if(err!=nil) {
		fmt.Println(err)
	}
	req.SetBasicAuth(ApiKey, "")
	cli := &http.Client{}
	resp, err := cli.Do(req)
	defer resp.Body.Close()
	if(err!=nil) {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	log.Printf("RESPONSE CODE: %d\n",resp.StatusCode)
	var errCheck models.TrackerList
	err = json.Unmarshal(body, &errCheck)
	if(err!=nil) {
		fmt.Println(err)
	}
	fmt.Println(errCheck)
	s := string(body[:])
	j := "JSON"
j += s
	fmt.Println(j)
}