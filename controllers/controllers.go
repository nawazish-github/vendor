package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/nawazish-github/vendor/models"
)

func VendorHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		createNewVendor(w, r.Body)
	}
}

func createNewVendor(w http.ResponseWriter, b io.ReadCloser) {
	decoder := json.NewDecoder(b)
	var v models.Vendor
	err := decoder.Decode(&v)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = v.ValidateVendorData()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
}
