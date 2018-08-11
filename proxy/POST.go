/**
* Copyright © 2017, ACM@UIUC
*
* This file is part of the Groot Project.
*
* The Groot Project is open source software, released under the University of
* Illinois/NCSA Open Source License. You should have received a copy of
* this license in a file with the distribution.
**/

package proxy

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/arbor-dev/arbor/logger"
)

// POST provides a proxy POST request allowing authorized clients to make POST
// requests of the microservices
//
// Pass the http Request from the client and the ResponseWriter it expects.
//
// Pass the target url of the backend service (not the url the client called).
//
// Pass the format of the service.
//
// Pass a authorization token (optional).
//
// Will call the service and return the result to the client.
func POST(w http.ResponseWriter, url string, format string, token string, r *http.Request) {

	preprocessing_err := requestPreprocessing(w, r)
	if preprocessing_err != nil {
		return
	}

	if format != "XML" && format != "JSON" { //TODO: Support Post form data
		err := errors.New("ERROR: unsupported data encoding")
		invalidPOST(w, err)
		logger.Log(logger.ERR, err.Error())
		return
	}
	contents, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		invalidPOST(w, err)
		logger.Log(logger.ERR, err.Error())
		return
	}
	if err := r.Body.Close(); err != nil {
		invalidPOST(w, err)
		logger.Log(logger.ERR, err.Error())
		return
	}

	origin := r.Header.Get("Origin")

	//TODO: FIGURE OUT ORIGIN RULES
	if origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}

	switch format {
	case "XML":
		xmlPOST(r, w, url, token, contents)
		return
	case "JSON":
		jsonPOST(r, w, url, token, contents)
		return
	default:
		invalidPOST(w, err)
		logger.Log(logger.ERR, "Unsupported Data Encoding")
		return
	}
}

func jsonPOST(r *http.Request, w http.ResponseWriter, url string, token string, contents []byte) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(contents))
	req.Header.Set("Content-Type", JSONHeader)
	req.Header.Set("Accept", "application/json")

	for k, vs := range r.Header {
		req.Header[k] = make([]string, len(vs))
		copy(req.Header[k], vs)
	}
	if token != "" {
		req.Header.Set("Authorization", token)
	}

	client := &http.Client{
		Timeout: time.Duration(Timeout) * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	resp, err := client.Do(req)
	logger.LogResp(logger.DEBUG, resp)

	if err != nil {
		if resp != nil {
			log.Println(resp.StatusCode)
		}
		invalidPOST(w, err)
		logger.Log(logger.ERR, err.Error())
		return
	} else if resp.StatusCode == http.StatusFound {
		logger.Log(logger.DEBUG, "Service Returned Redirect")
		w.Header().Set("Location", resp.Header.Get("Location"))
		w.WriteHeader(http.StatusFound)
		return
	} else if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		logger.Log(logger.WARN, "SERVICE FAILED - SERVICE RETURNED STATUS "+http.StatusText(resp.StatusCode))
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(resp.StatusCode)
	}

	defer resp.Body.Close()

	contents, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		invalidPOST(w, err)
		logger.Log(logger.ERR, "Failed to read response")
		return
	}

	var serverData interface{}
	err = json.Unmarshal(contents, &serverData)
	if err != nil {
		invalidPOST(w, err)
		logger.Log(logger.ERR, "Failed to unmarshal json "+err.Error())
		return
	}

	w.Header().Set("Content-Type", JSONHeader)

	if err := json.NewEncoder(w).Encode(serverData); err != nil {
		invalidPOST(w, err)
		logger.Log(logger.ERR, err.Error())
		return
	}

	//NOTE: Apparently not needed but add back in if things break
	//w.WriteHeader(http.StatusCreated)
}

func xmlPOST(r *http.Request, w http.ResponseWriter, url string, token string, contents []byte) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(contents))
	req.Header.Set("Content-Type", XMLHeader)

	for k, vs := range r.Header {
		req.Header[k] = make([]string, len(vs))
		copy(req.Header[k], vs)
	}
	if token != "" {
		req.Header.Set("Authorization", token)
	}

	client := &http.Client{
		Timeout: time.Duration(Timeout) * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	resp, err := client.Do(req)
	logger.LogResp(logger.DEBUG, resp)

	if err != nil {
		if resp != nil {
			log.Println(resp.StatusCode)
		}
		invalidPOST(w, err)
		logger.Log(logger.ERR, err.Error())
		return
	} else if resp.StatusCode == http.StatusFound {
		logger.Log(logger.DEBUG, "Service Returned Redirect")
		w.Header().Set("Location", resp.Header.Get("Location"))
		w.WriteHeader(http.StatusFound)
		return
	} else if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		logger.Log(logger.WARN, "SERVICE FAILED - SERVICE RETURNED STATUS "+http.StatusText(resp.StatusCode))
		w.Header().Set("Content-Type", "<?xml version=\"1.0\" encoding=\"UTF-8\"?>")
		w.WriteHeader(resp.StatusCode)
	}

	defer resp.Body.Close()

	contents, err = ioutil.ReadAll(resp.Body)
	var serverData interface{}
	err = xml.Unmarshal(contents, &serverData)
	if err != nil {
		invalidPOST(w, err)
		logger.Log(logger.ERR, err.Error())
		return
	}

	w.Header().Set("Content-Type", XMLHeader)

	if err := json.NewEncoder(w).Encode(serverData); err != nil {
		invalidPOST(w, err)
		logger.Log(logger.ERR, err.Error())
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func invalidPOST(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusUnprocessableEntity)
	data := map[string]interface{}{"Code": http.StatusUnprocessableEntity, "Text": "Unprocessable Entity", "Specfically": err}
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
