package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"google.golang.org/api/oauth2/v2"
)

var httpClient = &http.Client{}

var validDomains map[string]bool
var validEmails map[string]bool

func init() {
	validDomains = map[string]bool{}
	validEmails = map[string]bool{}
	value, exists := os.LookupEnv("VALID_DOMAINS")
	if exists {
		for _, d := range strings.Split(value, ",") {
			validDomains[d] = true
		}
	}

	value, exists = os.LookupEnv("VALID_EMAILS")
	if exists {
		for _, d := range strings.Split(value, ",") {
			validEmails[d] = true
		}
	}

	if len(validEmails)+len(validDomains) == 0 {
		panic("You must set either VALID_DOMAINS or VALID_EMAILS environment variables with a comma delimited list")
	}
}

// AuthMessage is a simple response to the browser
type AuthMessage struct {
	IDToken string `json:"token"`
	Next    string `json:"next"`
}

func verify(w http.ResponseWriter, r *http.Request) {
	var t AuthMessage
	next := r.URL.Query().Get("next")
	fmt.Println(next)
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		fmt.Println("Token body failed")
		fmt.Fprint(w, "{\"success\": false}")
		return
	}

	tokenInfo, err := verifyIDToken(t.IDToken)
	if err != nil {
		fmt.Printf("Token auth failed: %s\n", err)
		fmt.Fprint(w, "{\"success\": false}")
		return
	}
	w.Header().Set("Content-Type", "application/json")

	domain := strings.Split(tokenInfo.Email, "@")[1]

	if _, ok := validDomains[domain]; ok {
		setCookie(w, tokenInfo.Email)
		// note - the "next" part of this reply is not used, as the browser extracts from the current window location
		fmt.Fprintf(w, "{\"success\": true, \"next\": \"%s\"}", next)
		return
	}

	if _, ok := validEmails[tokenInfo.Email]; ok {
		setCookie(w, tokenInfo.Email)
		fmt.Fprintf(w, "{\"success\": true, \"next\": \"%s\"}", next)
		return
	}

	fmt.Fprint(w, "{\"success\": false}")
	return

}

func verifyIDToken(idToken string) (*oauth2.Tokeninfo, error) {
	oauth2Service, err := oauth2.New(httpClient)
	tokenInfoCall := oauth2Service.Tokeninfo()
	tokenInfoCall.IdToken(idToken)
	tokenInfo, err := tokenInfoCall.Do()
	if err != nil {
		return nil, err
	}
	return tokenInfo, nil
}
