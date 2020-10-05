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
		w.WriteHeader(400)
		return
	}
	var allowed bool
	if data.ISO {
		allowed, err = IPW.IsIPWhitelistedByISO(data.IP, data.Whitelist)
	} else {
		if data.Locale == "" {
			data.Locale = "en"
		}
		allowed, err = IPW.IsIPWhitelistedByLocale(data.IP, data.Locale, data.Whitelist)
	}
	var errorString string
	if err != nil {
		errorString = err.Error()
	}
	returnData := IPWhitelistReturn{
		allowed,
		errorString,
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
	IPInWhitelist bool 		`json:"ip_is_whitelisted"`
	Error string			`json:"error"`

}
