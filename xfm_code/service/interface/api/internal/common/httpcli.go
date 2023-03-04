package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

//基本的Get请求

func GetMethod(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	log.Println(string(body))
	log.Println(resp.StatusCode)
	if resp.StatusCode == 200 {
		log.Println("ok")
	}
}

//带参数的Get请求

func GetMethodOfMeter(url, request string) {
	resp, err := http.Get(url + request)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(body))
}

//参数为变量Get请求

func GetMethodOfValue(urls string) {
	params := url.Values{}
	Url, err := url.Parse(urls)
	if err != nil {
		return
	}
	params.Set("", "")
	params.Set("", "")
	//如果参数中有中文参数,这个方法会进行URLEncode
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	log.Println(urlPath)
	resp, err := http.Get(urlPath)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(body))
}

//基本的post请求

func PostMethod(urls string) {
	urlValues := url.Values{}
	urlValues.Add("", "")
	urlValues.Add("", "")
	resp, _ := http.PostForm(urls, urlValues)
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(body))
}

//带JSON参数的Post请求

func PostJsonMet(urls string) {
	client := &http.Client{}
	data := make(map[string]interface{})
	data[""] = ""
	data[""] = ""
	bytesData, _ := json.Marshal(data)
	req, _ := http.NewRequest("POST", urls, bytes.NewReader(bytesData))
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(body))
}
