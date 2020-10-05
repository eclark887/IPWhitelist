#IP Whitelist server
This project serves an API that takes in an IP and a whitelist of countries and will return if the IP is whitelisted.  

#How to setup locally (Without docker)
* golang 1.15
* make setup
* make ensure
* make build

# Local development
* make test - to run the unit & integration tests
* make fmt - to run the linting

#How to setup locally (With Docker)
* docker-compose up -d ip_whitelist_service
* It will take a minute to download the deps
* The server can be accessed on http://localhost:8080

#Update protobufs
* protoc -I IPWhitelist/ IPWhitelist/whitelist.proto --go_out=plugins=grpc:IPWhitelist

#API
POST /ipwhitelist
Content-type application/json
{
Locale (string, optional)
ISO (bool, optional)
IP (string, required)
whitelist (string[])
}
If you wish to use the ISO country codes instead of locale pass in "iso": true.
A locale can be of the following types ["en", "de", "es", "fr", "ja"] if no locale or is passed in then the API will default to "en"

The API will return a 200 & "ip_is_whitelisted": bool on a successful request.

#Automatic update of the GeoIP Country DB
The following env vars will need to be set: 
* GEOIPUPDATE_FREQUENCY - 24 (the db will be updated daily)
* GEOIPUPDATE_ACCOUNT_ID - Your MaxMind account ID.
* GEOIPUPDATE_LICENSE_KEY - Your case-sensitive MaxMind license key.
* GEOIPUPDATE_EDITION_IDS - GeoIP2-Country
* GEOIPUPDATE_PRESERVE_FILE_TIMES - 1 (Turns on saving the file update times)
* GEOIPUPDATE_VERBOSE - true
* docker run --env-file <file> -v packages/IPWhitelist:/usr/share/GeoIP maxmindinc/geoipupdate

Source: https://hub.docker.com/r/maxmindinc/geoipupdate

# Future updates
* Adding Prometheus to scrape metrics (response times, apdex)
* Adding an integration for a CI/CD platform (Semaphore, Jenkins, CircleCI to run the unit tests on PRs)
* Postman team collection & postman tests. Automation tests for staging to be ran before any releases to prod
* Documentation on Confluence or whatever internal documentation platform the company uses
* Finish the addition of the gRPC integration.  The function RegisterWhitelistServiceHandler in the IPWhitelist package needs to be filled out