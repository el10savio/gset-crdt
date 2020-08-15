package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// Lookup is the HTTP handler used to return 
// if a given value is present in the GSet node in the server
func Lookup(w http.ResponseWriter, r *http.Request) {
	var err error
	var present bool

	// Obtain the value from URL params
	value := mux.Vars(r)["value"]

	// Sync the GSets if multiple nodes
	// are present in a cluster 
	if len(GetPeerList()) != 0 {
		GSet, _ = Sync(GSet)
	}

	// Lookup given value in the GSet
	present, err = GSet.Lookup(value)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("failed to lookup gset value")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// DEBUG log in the case of success indicating
	// the new GSet, the lookup value and if its present
	log.WithFields(log.Fields{
		"set":     GSet.Set,
		"value":   value,
		"present": present,
	}).Debug("successful gset lookup")

	// Return HTTP 404 Not Found 
	// if the value is not present
	if present == false {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Return HTTP 200 OK if
	// thevalue is present
	w.WriteHeader(http.StatusOK)
}
