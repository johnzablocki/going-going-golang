package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
)

//TCP runs a TCP command
func TCP() {
	conn, err := net.Dial("tcp", "johnzablocki.com:80")
	Check(err)

	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')

	fmt.Println(status)
}

//CNAME gets a CNAME
func CNAME() {
	cname, err := net.LookupCNAME("checkout.rideopenroad.com")
	Check(err)

	fmt.Println(cname)
}

//Host gets a host
func Host() {
	host, err := net.LookupHost("checkout.rideopenroad.com")
	Check(err)

	fmt.Println(host)
}

//Http gets a result
func Http() {
	resp, err := http.Get("http://www.johnzablocki.com")
	Check(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	Check(err)
	fmt.Println(string(body))
}

func HttpClient() {

	tr := &http.Transport{
		DisableCompression: true,
		MaxConnsPerHost:    10,
	}

	client := http.Client{Transport: tr}
	resp, err := client.Get("http://www.johnzablocki.com")
	Check(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	Check(err)
	fmt.Println(string(body))
}
