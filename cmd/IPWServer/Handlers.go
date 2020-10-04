package main

import (
	IPW "ipw/packages/IPWhitelist"
	"encoding/json"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

// k8s healthcheck
func healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func postIPWhitelist(w http.ResponseWriter, r *http.Request) {
	// TODO error checking for incorrect additons to the API
	decoder := json.NewDecoder(r.Body)
	var data postData
	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}
	var allowed bool
	if data.ISO {
		allowed, err = IPW.IsIPWhitelistedByISO(data.IP, data.Whitelist)
	} else {
		allowed, err = IPW.IsIPWhitelistedByLocale(data.IP, data.Locale, data.Whitelist)
	}

	returnData := IPWhitelistReturn{
		allowed,
		err,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(returnData)
}

type postData struct {
	Locale    string
	ISO       bool
	IP        string
	Whitelist []string
}

type IPWhitelistReturn struct {
	IPInWhitelist bool 	`json:"ip_in_whitelist"`
	err error			`json:"error"`

}
