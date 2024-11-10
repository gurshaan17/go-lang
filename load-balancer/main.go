package main

import (
	"fmt"
	"net/http/httputil"
	"net/url"
	"os"
)

type simpleServer struct{
	addr	string
	proxy   *httputil.ReverseProxy
}

func newSimpleServer(addr string) *simpleServer{
	serverURL, err := url.Parse(addr)
	handleErr(err)
	return &simpleServer{
		addr: addr,
		proxy: httputil.NewSingleHostReverseProxy(serverURL),
	}
}

func handleErr(err error){
	if err != nil{
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
}

func main(){

}