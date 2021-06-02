package main

import (
	"net/http"
)

// Microservice tutorial_1
func main() {
	http.ListenAndServe(":9090", nil)
	//===============================================================================
	// this is a function of http binding an ip address
	// http.ListenAndServe() function has two parameter
	// First one is ip address port no to bind all the ip address
	// Second parameter is Handler 
	//===============================================================================
	


}
