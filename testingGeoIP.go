package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var print = fmt.Println

// struct geoResponse {

// 		"ip": "8.8.8.8",
// 		"hostname": "google-public-dns-a.google.com",
// 		"continent_code": "NA",
// 		"continent_name": "North America",
// 		"country_code2": "US",
// 		"country_code3": "USA",
// 		"country_name": "United States",
// 		"country_capital": "Washington",
// 		"state_prov": "California",
// 		"district": "",
// 		"city": "Mountain View",
// 		"zipcode": "94043",
// 		"latitude": "37.4229",
// 		"longitude": "-122.085",
// 		"is_eu": false,
// 		"calling_code": "+1",
// 		"country_tld": ".us",
// 		"languages": "en-US,es-US,haw,fr",
// 		"country_flag": "https://ipgeolocation.io/static/flags/us_64.png",
// 		"isp": "Level 3 Communications",
// 		"connection_type": "",
// 		"organization": "Google Inc.",
// 		"geoname_id": "5375480",
// 		"currency": {
// 			"code": "USD",
// 			"name": "US Dollar",
// 			"symbol": "$"
// 		},
// 		"time_zone": {
// 			"name": "America/Los_Angeles",
// 			"offset": -8,
// 			"current_time": "2019-01-14 03:30:00.135-0800",
// 			"current_time_unix": 1547465400.135,
// 			"is_dst": false,
// 			"dst_savings": 1
// 		}

// }

func getGeoInfo(ipaddress string) string {
	var address = ipaddress
	var API_Key = "4d1c81287462457797aa070514b60faa"

	temp := fmt.Sprintf("https://api.ipgeolocation.io/ipgeo?apiKey=%s&ip=%s", API_Key, address)

	resp, err := http.Get(temp)
	if err != nil {
		// handle error

	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	return string(body)

}

func main() {
	// Provide a domain name or IP address

	print(getGeoInfo("45.79.8.237"))
}
