package goeasypost

import (
	"flag"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	flag.Parse()
	//	DisableLogger()
	os.Exit(m.Run())
}

func TestListTrackers(t *testing.T) {

	_, err := GetTrackerList()
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	//data, _ := json.MarshalIndent(trackList, "", "    ")
	//t.Log(string(data[:]))

}

func TestInvalidCreateTracker(t *testing.T) {

	_, err := NewTracker("EZ1000000001c", "USPS")
	if err == nil {
		t.Log(err)
		t.Fail()
	}
}

func TestValidCreateTracker(t *testing.T) {

	tracker, err := NewTracker("EZ4000000004", "UPS")
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	//	data, _ := json.MarshalIndent(tracker, "", "    ")
	//	t.Log(string(data[:]))
	if "delivered" != tracker.Status {
		t.Log(fmt.Sprintf("Expected status 'delivered', saw status %v", tracker.Status))
		t.Fail()
	}
}

func TestFlattenStruct(t *testing.T) {

	var a = new(Address)
	var v = new(Verifications)
	var vdelivery = new(Verification)

	a.Id = "xxx"
	a.Name = "Testing Tester"
	a.Verifications = v
	v.Delivery = vdelivery
	vdelivery.Success = true


	params := make(map[string]string)

	err := flattenStructMap(a, "address", params)
	if err != nil {
		t.Log(fmt.Sprintf("Expected error value nil, saw error %v", err))
		t.Fail()
	}

	for k, v := range params {
		fmt.Printf("%v=%v\n", k, v)
	}

	//t.Log(fmt.Sprintf("%v",params))
}

func TestNewAddress(t *testing.T) {
	var a = new(Address)

	//a.Id="xxx"
	a.Name = "Testing Tester"
	a.Street1 = "18495 lakeview lane"
	a.City = "Mount Vernon"
	a.State = "WA"
	a.Zip = "984"
	a.Country = "US"
	v := VerifyRequest{true,"delivery"}
	a.VerifyRequest = &v

	newAddr, err := NewAddress(a)
	if err != nil {
		t.Log(fmt.Sprintf("Expected error value nil, saw error %v", err))
		t.Fail()
	}

	spew.Dump(newAddr)

	retrievedAddr, err := RetrieveAddress(newAddr.Id)
	if err != nil {
		t.Log(fmt.Sprintf("Expected error value nil, saw error %v", err))
		t.Fail()
	}

	if(retrievedAddr.Id != newAddr.Id) {
		t.Log(fmt.Sprintf("Expected matching address ids, saw created id %v, retrieved id %v", newAddr.Id,retrievedAddr.Id))
		t.Fail()
	}

}


func TestNewParcel(t *testing.T) {
	var a = new(Parcel)

	//a.Id="xxx"
	a.Weight = 1.50
	a.Length = 6
	a.Height = 6
	a.Width = 6


	newParcel, err := NewParcel(a)
	if err != nil {
		t.Log(fmt.Sprintf("Expected error value nil, saw error %v", err))
		t.Fail()
	}

	spew.Dump(newParcel)

	retrievedParcel, err := RetrieveParcel(newParcel.Id)
	if err != nil {
		t.Log(fmt.Sprintf("Expected error value nil, saw error %v", err))
		t.Fail()
	}

	if(retrievedParcel.Id != newParcel.Id) {
		t.Log(fmt.Sprintf("Expected matching address ids, saw created id %v, retrieved id %v", newParcel.Id,retrievedParcel.Id))
		t.Fail()
	}

}


