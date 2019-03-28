package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

var print = fmt.Println
var api_key = "4d1c81287462457797aa070514b60faa"

type geoResponse struct {
	IP             string `json:"ip"`
	ContinentCode  string `json:"continent_code"`
	ContinentName  string `json:"continent_name"`
	CountryCode2   string `json:"country_code2"`
	CountryCode3   string `json:"country_code3"`
	CountryName    string `json:"country_name"`
	CountryCapital string `json:"country_capital"`
	StateProv      string `json:"state_prov"`
	District       string `json:"district"`
	City           string `json:"city"`
	Zipcode        string `json:"zipcode"`
	Latitude       string `json:"latitude"`
	Longitude      string `json:"longitude"`
	IsEu           bool   `json:"is_eu"`
	CallingCode    string `json:"calling_code"`
	CountryTld     string `json:"country_tld"`
	Languages      string `json:"languages"`
	CountryFlag    string `json:"country_flag"`
	Isp            string `json:"isp"`
	ConnectionType string `json:"connection_type"`
	Organization   string `json:"organization"`
	GeonameID      string `json:"geoname_id"`
	Currency       struct {
		Code   string `json:"code"`
		Name   string `json:"name"`
		Symbol string `json:"symbol"`
	} `json:"currency"`
	TimeZone struct {
		Name            string  `json:"name"`
		Offset          int     `json:"offset"`
		CurrentTime     string  `json:"current_time"`
		CurrentTimeUnix float64 `json:"current_time_unix"`
		IsDst           bool    `json:"is_dst"`
		DstSavings      int     `json:"dst_savings"`
	} `json:"time_zone"`
}

func getGeoInfo(ipaddress string) geoResponse {
	var address = ipaddress

	temp := fmt.Sprintf("https://api.ipgeolocation.io/ipgeo?apiKey=%s&ip=%s", api_key, address)

	resp, err := http.Get(temp)
	if err != nil {
		// handle error

	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	ret := geoResponse{}
	json.Unmarshal(body, &ret)
	return ret

}

func getLocaleData(ip string) (string, string, float32, float32) {
	ret := getGeoInfo(ip)
	lat, _ := strconv.ParseFloat(ret.Latitude, 32)
	long, _ := strconv.ParseFloat(ret.Longitude, 32)

	return ret.CountryName, ret.City, float32(lat), float32(long)
}

// func main() {
// 	// Provide a domain name or IP address

// 	print(getGeoInfo("45.79.8.237"))
// }
