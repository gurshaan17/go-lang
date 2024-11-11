package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

type Server interface{
	Address() string
	IsAlive() bool
	Serve(w http.ResponseWriter, r *http.Request) 
}

type SimpleServer struct{
	addr	string
	proxy   *httputil.ReverseProxy
}

func newSimpleServer(addr string) *SimpleServer{
	serverURL, err := url.Parse(addr)
	handleErr(err)
	return &SimpleServer{
		addr: addr,
		proxy: httputil.NewSingleHostReverseProxy(serverURL),
	}
}

type LoadBalancer struct{
	port 				string
	roundRobinCount     int
	servers				[]Server	
}

func NewLoadBalancer(port string, servers []Server) *LoadBalancer{
	return &LoadBalancer{
		port: port,
		roundRobinCount: 0,
		servers: servers,
	}
}

func handleErr(err error){
	if err != nil{
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
}

func (s *SimpleServer) Address() string {
	return s.addr
}

func (s *SimpleServer) IsAlive() bool {
	return true
}

func (s *SimpleServer) Serve(w http.ResponseWriter, r *http.Request) {
	s.proxy.ServeHTTP(w, r)
}

func (lb *LoadBalancer) getNextAvailableServer() Server{
	server := lb.servers[lb.roundRobinCount%len(lb.servers)]
	for !server.IsAlive() {
		lb.roundRobinCount++
		server = lb.servers[lb.roundRobinCount%len(lb.servers)]
	}
	lb.roundRobinCount++
	return server
}

func (lb *LoadBalancer) serveProxy(w http.ResponseWriter, r *http.Request){
	targetServer := lb.getNextAvailableServer()
	fmt.Printf("forwarding request to address %q\n", targetServer.Address())
	targetServer.Serve(w, r)
}

func main(){
	servers := []Server{
		newSimpleServer("https://facebook.com"),
		newSimpleServer("https://bing.com"),
		newSimpleServer("https://instagram.com"),
	}
	lb := NewLoadBalancer("8000", servers)
	handleRedirect := func (w http.ResponseWriter, r *http.Request){
		lb.serveProxy(w, r)
	}
	http.HandleFunc("/", handleRedirect)

	fmt.Printf("Server is serving requests at 'localhost:%s'\n",lb.port)

	http.ListenAndServe(":"+lb.port, nil)
}