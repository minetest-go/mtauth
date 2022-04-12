package web

import (
	"fmt"
	"net/http"
)

func NewOAuthController() *OAuthController {
	return &OAuthController{}
}

type OAuthController struct{}

func (ac *OAuthController) Authorize(resp http.ResponseWriter, req *http.Request) {
	q := req.URL.Query()
	client_id := q.Get("client_id")
	redirect_uri := q.Get("redirect_uri")
	scope := q.Get("scope")
	response_type := q.Get("response_type")
	response_mode := q.Get("response_mode")
	nonce := q.Get("nonce")

	fmt.Printf(
		"client_id: %s, redirect_uri: %s, scope: %s, response_type: %s, response_mode: %s, nonce: %s\n",
		client_id, redirect_uri, scope, response_type, response_mode, nonce,
	)

	http.Redirect(resp, req, "../../login", http.StatusFound)
}
