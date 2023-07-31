package handlers

import (
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://http.cat/images/420.jpg", http.StatusSeeOther)
}
