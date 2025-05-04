package helpers

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func PrintResponseHelper(resp *http.Response, debug bool) {
	if !debug {
		return
	}

	respByte, err := httputil.DumpResponse(resp, true)
	if err != nil {
		fmt.Println(fmt.Errorf("httputil.DumpResponse: %w", err))
		return
	}
	fmt.Println("============== Response")
	fmt.Println(string(respByte))
	fmt.Println("==============")
}

func PrintRequestHelper(req *http.Request, debug bool) {
	if !debug {
		return
	}

	reqByte, err := httputil.DumpRequest(req, true)
	if err != nil {
		fmt.Println(fmt.Errorf("httputil.DumpRequest: %w", err))
	}
	fmt.Println("============== Request")
	fmt.Println(string(reqByte))
	fmt.Println("==============")
}
