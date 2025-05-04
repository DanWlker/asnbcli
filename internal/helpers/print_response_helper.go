package helpers

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func PrintResponseHelper(resp *http.Response) {
	respByte, err := httputil.DumpResponse(resp, true)
	if err != nil {
		DebugLogger.Println(fmt.Errorf("httputil.DumpResponse: %w", err))
		return
	}
	DebugLogger.Println("============== Response")
	DebugLogger.Println(string(respByte))
	DebugLogger.Println("==============")
}

func PrintRequestHelper(req *http.Request) {
	reqByte, err := httputil.DumpRequest(req, true)
	if err != nil {
		DebugLogger.Println(fmt.Errorf("httputil.DumpRequest: %w", err))
	}
	DebugLogger.Println("============== Request")
	DebugLogger.Println(string(reqByte))
	DebugLogger.Println("==============")
}
