package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// Append ...
func Append(w http.ResponseWriter, r *http.Request) {
	var err error

	// Obtain the value from URL params
	value := mux.Vars(r)["value"]

	GSet.Set, err = GSet.Append(value)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("failed to append value")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.WithFields(log.Fields{
		"set":   GSet.Set,
		"value": value,
	}).Debug("successful gset append")

	w.WriteHeader(http.StatusOK)
}
