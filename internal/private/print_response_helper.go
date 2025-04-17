package private

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func PrintResponseHelper(resp *http.Response) {
	respByte, err := httputil.DumpResponse(resp, true)
	if err != nil {
		fmt.Println(fmt.Errorf("httputil.DumpResponse: %w", err))
		return
	}
	fmt.Println("")
	fmt.Println(string(respByte))
	fmt.Println("")
}

func PrintRequestHelper(req *http.Request) {
	reqByte, err := httputil.DumpRequest(req, true)
	if err != nil {
		fmt.Println(fmt.Errorf("httputil.DumpRequest: %w", err))
	}
	fmt.Println("")
	fmt.Println(string(reqByte))
	fmt.Println("")
}
