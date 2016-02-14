package goeasypost
import (
	"fmt"
	"flag"
	"net/http"
)


func ExampleMain() {
	flag.Parse()
	fmt.Println("hello main" )
}

func ExampleGet() {

	req, err := http.NewRequest("GET", "https://api.easypost.com/v2", nil)
	if(err!=nil) {
		fmt.Println(err)
	}
	req.SetBasicAuth(ApiKey, "")
	cli := &http.Client{}
	resp, err := cli.Do(req)
	    if(err!=nil) {
			fmt.Println(err)
		}
	fmt.Println(resp)
}
