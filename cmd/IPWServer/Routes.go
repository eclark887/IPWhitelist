package IPWServer

import (
	"log"
	"net/http"
)

func handleRequests() {
	http.HandleFunc("/", homePage)
	// k8s health checks
	http.HandleFunc("/healthz", healthz)
	http.HandleFunc("/livez", healthz)
	http.HandleFunc("/readyz", healthz)
	log.Fatal(http.ListenAndServe(":1000", nil))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

// k8s healthcheck
func healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func main() {
	handleRequests()
}