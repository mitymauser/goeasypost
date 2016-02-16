package goeasypost

import (
	"testing"
	"os"
	"flag"
	"fmt"
)


func TestMain(m *testing.M) {
	flag.Parse()
	DisableLogger()
	os.Exit(m.Run())
}


func TestInvalidCreateTracker(t *testing.T) {

	_, err := CreateTracker("EZ1000000001c", "USPS")
	if (err == nil) {
		t.Log(err)
		t.Fail()
	}
}


func TestValidCreateTracker(t *testing.T) {

	tracker, err := CreateTracker("EZ4000000004", "UPS")
	if (err != nil) {
		t.Log(err)
		t.Fail()
	}

	if("delivered" != tracker.Status) {
		t.Log(fmt.Sprintf("Expected status 'delivered', saw status %v",tracker.Status))
		t.Fail()
	}
}
