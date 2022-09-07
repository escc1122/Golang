package main

import (
	"encoding/json"
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
	req, _ := http.NewRequest(http.MethodGet, "https://www.google.com.tw", nil)
	req.Header.Set("User-Agent", "Mozilla/5.0")
	//req.Header.Set("Content-Type", "application/json")
	//req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
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

func noRedirect() {
	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodGet, "https://www.google.com.tw", nil)
	//noRedirect
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}
	res, err := client.Do(req)

	if err != nil {
		fmt.Println("StatusCode", res.StatusCode)
		if res.StatusCode == 302 {
			fmt.Println("got redirect")
		}
	}
	defer res.Body.Close()
	fmt.Println("Http status code:", res.StatusCode)
	sitemap, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", sitemap)
}

type githubData struct {
	CurrentUserUrl string `json:"current_user_url"`
}

var githubDataInterFace interface{}

func withJson() {
	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodGet, "https://api.github.com", nil)
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
	githubData := githubData{}

	err = json.Unmarshal(sitemap, &githubData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", githubData.CurrentUserUrl+"\n")

	githubDataAll := githubDataInterFace
	err = json.Unmarshal(sitemap, &githubDataAll)
	if err != nil {
		log.Fatal(err)
	}

	m, ok := githubDataAll.(map[string]interface{})
	fmt.Println(ok)
	fmt.Printf("%s", m["current_user_url"])

}

func withCookie() {
	var cookie = &http.Cookie{
		Name:   "cookie_name",
		Value:  "cookie_value",
		MaxAge: 300,
	}
	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodGet, "https://www.google.com.tw", nil)
	req.AddCookie(cookie)
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
	//withHeader()
	withJson()
	//noRedirect()
	//withCookie()
}
