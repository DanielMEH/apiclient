package routers

import "net/http"

func HomeHandle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World mundo!!"))
}
