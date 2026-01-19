package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// In a real app, these would come from a database
var users = map[string]string{
	"alice": "password123",
	"bob":   "secret456",
}

// Map of user -> topic -> permissions (read/write)
var acls = map[string]map[string]string{
	"alice": {
		"sensors/alice": "rw",
		"sensors/#":     "r",
	},
	"bob": {
		"sensors/bob": "rw",
	},
}

func main() {
	http.HandleFunc("/user", handleUser)
	http.HandleFunc("/superuser", handleSuperuser)
	http.HandleFunc("/acl", handleACL)

	fmt.Println("Auth Service starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleUser(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	fmt.Printf("[Auth] Body: %s\n", string(body))

	var req AuthRequest
	// Try parsing as JSON
	if err := json.Unmarshal(body, &req); err != nil {
		// If not JSON, try Form values
		r.ParseForm()
		req.Username = r.FormValue("username")
		req.Password = r.FormValue("password")
	}

	username := req.Username
	password := req.Password

	fmt.Printf("[Auth] Login attempt for user: %s\n", username)

	if expectedPass, ok := users[username]; ok && expectedPass == password {
		w.WriteHeader(http.StatusOK)
		return
	}

	w.WriteHeader(http.StatusUnauthorized)
}

func handleSuperuser(w http.ResponseWriter, r *http.Request) {
	// For this lab, no superusers
	w.WriteHeader(http.StatusUnauthorized)
}

type ACLRequest struct {
	Username string `json:"username"`
	Topic    string `json:"topic"`
	Acc      int    `json:"acc"`
}

func handleACL(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	fmt.Printf("[ACL] Body: %s\n", string(body))

	var req ACLRequest
	if err := json.Unmarshal(body, &req); err != nil {
		r.ParseForm()
		req.Username = r.FormValue("username")
		req.Topic = r.FormValue("topic")
		// For form values, Acc might be a string
	}

	username := req.Username
	topic := req.Topic
	acc := req.Acc

	fmt.Printf("[ACL] Access check for user: %s, topic: %s, access: %d\n", username, topic, acc)

	userAcls, ok := acls[username]
	if !ok {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Simple check (not handling wildcards properly for this basic lab)
	for t, perm := range userAcls {
		if t == topic || (strings.HasSuffix(t, "#") && strings.HasPrefix(topic, strings.TrimSuffix(t, "#"))) {
			// MOSQ_ACL_READ = 1, MOSQ_ACL_WRITE = 2, MOSQ_ACL_SUBSCRIBE = 4
			if (acc == 1 || acc == 4) && strings.Contains(perm, "r") {
				w.WriteHeader(http.StatusOK)
				return
			}
			if acc == 2 && strings.Contains(perm, "w") {
				w.WriteHeader(http.StatusOK)
				return
			}
		}
	}

	w.WriteHeader(http.StatusForbidden)
}
