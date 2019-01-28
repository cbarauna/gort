package service

import (
	"net/http"

	"github.com/gorilla/mux"
)

// handleDeleteBundleVersion handles "DELETE /v2/bundles/{bundlename}/versions/{version}"
func handleDeleteBundleVersion(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Implemented", http.StatusNotImplemented)
	return
}

// handleGetBundleVersion handles "GET /v2/bundles/{bundlename}/versions/{version}"
func handleGetBundleVersion(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Implemented", http.StatusNotImplemented)
	return
}

// handleGetBundles handles "GET /v2/bundles"
func handleGetBundles(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Implemented", http.StatusNotImplemented)
	return
}

// handleGetBundleVersions handles "GET /v2/bundles/{bundlename}"
func handleGetBundleVersions(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Implemented", http.StatusNotImplemented)
	return
}

// handlePostBundleVersion handles "POST /v2/bundles/{bundlename}/versions/{version}"
func handlePostBundleVersion(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Implemented", http.StatusNotImplemented)
	return
}

func addBundleMethodsToRouter(router *mux.Router) {
	router.HandleFunc("/v2/bundles", handleGetBundles).Methods("GET")
	router.HandleFunc("/v2/bundles/{bundlename}", handleGetBundleVersions).Methods("GET")
	router.HandleFunc("/v2/bundles/{bundlename}/versions/{version}", handleGetBundleVersion).Methods("GET")
	router.HandleFunc("/v2/bundles/{bundlename}/versions/{version}", handlePostBundleVersion).Methods("POST")
	router.HandleFunc("/v2/bundles/{bundlename}/versions/{version}", handleDeleteBundleVersion).Methods("DELETE")
}
