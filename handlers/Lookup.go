package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// Lookup ...
func Lookup(w http.ResponseWriter, r *http.Request) {
	var err error
	var present bool

	// Obtain the value from URL params
	value := mux.Vars(r)["value"]

	if len(GetPeerList()) != 0 {
		GSet, _ = Sync(GSet)
	}

	present, err = GSet.Lookup(value)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("failed to lookup gset value")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.WithFields(log.Fields{
		"set":     GSet.Set,
		"value":   value,
		"present": present,
	}).Debug("successful gset lookup")

	if present == false {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}
