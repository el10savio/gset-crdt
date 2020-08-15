package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

	"../gset"
)

// Sync ...
func Sync(GSet gset.GSet) (gset.GSet, error) {
	peers := GetPeerList()
	if len(peers) == 0 {
		return GSet, errors.New("nil peers present")
	}

	for _, peer := range peers {
		peerGset, err := SendListRequest(peer)
		if err != nil {
			log.WithFields(log.Fields{"error": err, "peer": peer}).Error("failed sending gset values request")
			continue
		}

		if len(peerGset.Set) == 0 {
			continue
		}

		GSet, _ = gset.Merge(GSet, peerGset)
	}

	log.WithFields(log.Fields{
		"set": GSet,
	}).Debug("successful gset sync")

	return GSet, nil
}

// SendListRequest ...
func SendListRequest(peer string) (gset.GSet, error) {
	var _gset gset.GSet

	if peer == "" {
		return _gset, errors.New("empty peer provided")
	}

	url := fmt.Sprintf("http://%s.%s/gset/values", peer, GetNetwork())

	response, err := SendRequest(url)
	if err != nil {
		return _gset, err
	}

	if response.StatusCode != http.StatusOK {
		return _gset, errors.New("received invalid http response status:" + string(response.StatusCode))
	}

	var values []string
	err = json.NewDecoder(response.Body).Decode(&values)
	if err != nil {
		return _gset, err
	}

	_gset.Set = values
	return _gset, nil
}
