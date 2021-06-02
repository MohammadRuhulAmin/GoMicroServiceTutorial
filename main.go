package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Microservice tutorial_1
func main() {

	// reading everithing to variable d
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello World")
		d, _ := ioutil.ReadAll(r.Body)
		log.Printf("Data %s\n", d)

	})

	http.HandleFunc("/responseWriter", func(rw http.ResponseWriter, r *http.Request) {
		d, _ := ioutil.ReadAll(r.Body)
		fmt.Fprintf(rw, "Hellow %s", d)
	})

	http.HandleFunc("/errorMessage", func(rw http.ResponseWriter, r *http.Request) {
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			//===================================================
			//http.Error(rw,"OPS",http.StatusBadRequest)
			//===================================================
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte("Opps.. bad Request !"))
			return 
			//================================================================= 
			// return the requst .. otherwise this request will not terminate!
			// WriteHeader is a request handler status code !
			//=================================================================
		}

		log.Printf("Data %s",d)

	})

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("GoodBye World")
	})

	//===============================================================================
	// this is a function of http binding an ip address
	// http.ListenAndServe() function has two parameter
	// First one is ip address port no to bind all the ip address
	// Second parameter is Handler
	// go other terminal and type curl -v localhost:port/link
	//===============================================================================

	http.ListenAndServe(":9090", nil)

}
