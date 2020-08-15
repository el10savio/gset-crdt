package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

	"../gset"
)

// Sync merges multiple GSets present in a network to get them in sync
// It does so by obtaining the GSet from each node in the cluster
// and performs a merge operation with the local GSet
func Sync(GSet gset.GSet) (gset.GSet, error) {
	// Obtain addresses of peer nodes in the cluster
	peers := GetPeerList()

	// Return the local GSet back if no peers
	// are present along with an error
	if len(peers) == 0 {
		return GSet, errors.New("nil peers present")
	}

	// Iterate over the peer list and send a /gset/values GET request 
	// to each peer to obtain its GSet
	for _, peer := range peers {
		peerGset, err := SendListRequest(peer)
		if err != nil {
			log.WithFields(log.Fields{"error": err, "peer": peer}).Error("failed sending gset values request")
			continue
		}

		// Skip merge if the peer's GSet is empty
		if len(peerGset.Set) == 0 {
			continue
		}

		// Merge the peer's GSet with our local GSet
		GSet, _ = gset.Merge(GSet, peerGset)
	}

	// DEBUG log in the case of success
	// indicating the new GSet
	log.WithFields(log.Fields{
		"set": GSet.Set,
	}).Debug("successful gset sync")

	// Return the synced new GSet
	return GSet, nil
}

// SendListRequest is used to send a GET /gset/values
// to peer nodes in the cluster
func SendListRequest(peer string) (gset.GSet, error) {
	var _gset gset.GSet

	// Return an empty GSet followed by an error if the peer is nil
	if peer == "" {
		return _gset, errors.New("empty peer provided")
	}

	// Resolve the Peer ID and network to generate the request URL
	url := fmt.Sprintf("http://%s.%s/gset/values", peer, GetNetwork())
	response, err := SendRequest(url)
	if err != nil {
		return _gset, err
	}

	// Return an empty GSet followed by an error 
	// if the peer's response is not HTTP 200 OK
	if response.StatusCode != http.StatusOK {
		return _gset, errors.New("received invalid http response status:" + string(response.StatusCode))
	}

	// Decode the peer's GSet to be usable by our local GSet
	var values []string
	err = json.NewDecoder(response.Body).Decode(&values)
	if err != nil {
		return _gset, err
	}

	// Return the decoded peer's GSet
	_gset.Set = values
	return _gset, nil
}
