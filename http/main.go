package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func base() {
	res, err := http.Get("https://www.google.com.tw")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	sitemap, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", sitemap)
}

func withHeader() {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://www.google.com.tw", nil)
	req.Header.Set("User-Agent", "Mozilla/5.0")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Http get err:", err)
	}
	defer res.Body.Close()
	fmt.Println("Http status code:", res.StatusCode)
	sitemap, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", sitemap)
}

func main() {
	withHeader()
}
