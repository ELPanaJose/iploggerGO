package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("server on")
	http.HandleFunc("/", getIP)

	port, ok := os.LookupEnv("PORT")

	if ok == false {
		port = "5000"
	}

	log.Println(http.ListenAndServe(":"+port, nil))
}

func getIP(w http.ResponseWriter, r *http.Request) {
	ip := r.Header.Get("x-forwarded-for")
	log.Println("new ip : " + ip)
	resp, _ := http.Get("http://ip-api.com/json/"+ip) //hace el request
	outJson, _ := ioutil.ReadAll(resp.Body)                 // decodifica el request
	var out bytes.Buffer                                    // es el out de la funcion de abajo
	json.Indent(&out, outJson, "", "\t")                    // imprime bonito el request
	fmt.Println(string(out.Bytes()))                        // imprime el request
	http.Redirect(w, r, "https://monda3213.net/", 301)
}
