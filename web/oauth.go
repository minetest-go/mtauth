package web

import (
	"fmt"
	"net/http"
	"net/url"
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

	// TODO: check if logged-in

	u, err := url.Parse("http://127.0.0.1:8080/login")
	if err != nil {
		SendError(resp, 500, err.Error())
		return
	}

	nq := u.Query()
	nq.Add("client_id", client_id)
	nq.Add("return_to", req.URL.String())
	u.RawQuery = nq.Encode()

	http.Redirect(resp, req, u.String(), http.StatusFound)
}
