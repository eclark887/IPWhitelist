package IPWhitelist

import (
	"errors"
	"log"
	"net"
	"os"

	geoip2 "github.com/oschwald/geoip2-golang"
)

var IPFormatErr = errors.New("the IP is incorrectly formatted")
var LocaleErr = errors.New("locale string is not allowed")

func GetRecordFromIP(IP string) (*geoip2.Country, error) {
	ip := net.ParseIP(IP)
	if ip == nil {
		return nil, IPFormatErr
	}
	// TODO move the open & close to separate functions so we can manage the connections & use a pool for connections to cut down overhead of opening the file
	rootdir := os.Getenv("REPOPATH")
	db, err := geoip2.Open(rootdir + "/packages/IPWhitelist/GeoLite2-Country.mmdb")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer db.Close()

	record, err := db.Country(ip)
	if err != nil {
		log.Println(err)
	}
	return record, err
}

func GetCountryNameFromRecord(record *geoip2.Country, locale string) (string, error) {
	if VerifyAllowedLocale(locale) {
		return record.Country.Names[locale], nil
	} else {
		return "", LocaleErr
	}
}

func GetISOFromRecord(record *geoip2.Country) string {
	return record.Country.IsoCode
}

func VerifyAllowedLocale(locale string) bool {
	locales := []string{"en", "de", "es", "fr", "ja"}
	return find(locales, locale)
}

func find(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func IsIPWhitelistedByLocale(ip string, locale string, countries []string) (bool, error) {
	if !VerifyAllowedLocale(locale) {
		return false, LocaleErr
	}
	record, err := GetRecordFromIP(ip)
	if err != nil {
		return false, err
	}
	country, err := GetCountryNameFromRecord(record, locale)
	if err != nil {
		return false, err
	}
	return find(countries, country), nil
}

func IsIPWhitelistedByISO(ip string, countries []string) (bool, error) {
	record, err := GetRecordFromIP(ip)
	if err != nil {
		return false, err
	}
	country := GetISOFromRecord(record)
	return find(countries, country), nil
}
