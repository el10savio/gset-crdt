package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// Append is the HTTP handler used to append
// values to the GSet node in the server
func Append(w http.ResponseWriter, r *http.Request) {
	var err error

	// Obtain the value from URL params
	value := mux.Vars(r)["value"]

	// Append the given value to our stored GSet
	GSet.Set, err = GSet.Append(value)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("failed to append value")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// DEBUG log in the case of success indicating
	// the new GSet and the value appended
	log.WithFields(log.Fields{
		"set":   GSet.Set,
		"value": value,
	}).Debug("successful gset append")

	// Return HTTP 200 OK in the case of success
	w.WriteHeader(http.StatusOK)
}
