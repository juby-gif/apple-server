package controllers

import (
	"net/http"
)

func (c *Controller) getVersion(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("apple-server-v1.0"))
}
