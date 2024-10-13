package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func LoadTube(w http.ResponseWriter, request *http.Request) {
	// get the basic document info from mongo and build the response object from the user data there
	// for now just return something that looks right

	response := map[string]interface{}{
		"@context": []string{
			"https://www.w3.org/ns/activitystreams",
			"https://w3id.org/security/v1",
		},
		"id":                "https://dontremember.me/tubes/angelsin",
		"type":              "Person",
		"preferredUsername": "angelsin",
		"inbox":             "https://dontremember.me/tubes/angelsin/inbox",
		"outbox":            "https://dontremember.me/tubes/angelsin/outbox",
		"endpoints": []map[string]string{
			{
				"sharedInbox": "https://bugle.lol/inbox",
			},
		},
		"name":                      "AngelSin",
		"summary":                   "Site owner",
		"url":                       "https://dontremember.me/tubes/angelsin",
		"manuallyApprovesFollowers": false,
		"discoverable":              true,
		"published":                 "2024-9-30T00:00:00Z",
		"publicKey": []map[string]string{
			{
				"id":           "https://dontremember.me/tubes/angelsin#main-key",
				"owner":        "https://dontremember.me/tubes/angelsin",
				"publicKeyPem": "-----BEGIN PUBLIC KEY-----...-----END PUBLIC KEY-----",
			},
		},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func WebfingerHandler(w http.ResponseWriter, r *http.Request) {
	//resource := r.URL.Query().Get("resource")
	// Generate JSON response based on the resource. todo: actually check the database to get this info
	// search db to see if acct (r.URL.Query().Get("resource")) exists
	// if it does then return the value if not return error
	response := map[string]interface{}{

		"subject": "acct:angelsin@dontremember.me",
		"links": []map[string]string{
			{
				"rel":  "self",
				"type": "application/activity+json",
				"href": "https://dontremember.me/tubes/angelsin",
			},
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func HandleInbox(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		// serve inbox?
		break
	case http.MethodPost:
		// add to inbox in db?
		break
	case http.MethodDelete:
		// remove from inbox
		break
	default:
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func HandleUserInbox(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		// serve inbox?
		break
	case http.MethodPost:
		// add to inbox in db?
		break
	case http.MethodDelete:
		// remove from inbox
		break
	default:
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func HandleOutbox(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		// serve outbox? read the public posts
		break
	case http.MethodPost:
		// add to outbox in db? post publicly
		break
	case http.MethodDelete:
		// remove from outbox. delete public post
		break
	default:
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func HandleCollections(writer http.ResponseWriter, request *http.Request) {

}
func HandleImg(writer http.ResponseWriter, request *http.Request) {

}

func HandleRoot(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("received root request")
	host := request.Host
	fmt.Println(host)
	var dir string
	if host == "portfolio.dontremember.me" {
		// Serve the directory for "portfolio.dontremember.me"
		dir = "./client/dist/portfolio"
	} else {
		// Serve a different directory for all other hosts
		dir = "./client/dist/main"
	}

	// Create a file server handler for the chosen directory
	fs := http.FileServer(http.Dir(dir))

	// Strip the prefix from the URL path and serve the files
	http.StripPrefix("/", fs).ServeHTTP(writer, request)

}
