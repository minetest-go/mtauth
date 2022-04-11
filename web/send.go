package web

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func SendError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	w.WriteHeader(code)
	w.Write([]byte(message))
}

func SendText(w http.ResponseWriter, txt string) {
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(txt))
}

func SendJson(w http.ResponseWriter, o interface{}) []byte {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	buf := bytes.NewBuffer([]byte{})
	json.NewEncoder(buf).Encode(o)
	w.Write(buf.Bytes())
	return buf.Bytes()
}

func Send(w http.ResponseWriter, o interface{}, err error) {
	if err != nil {
		SendError(w, 500, err.Error())
	} else {
		SendJson(w, o)
	}
}
