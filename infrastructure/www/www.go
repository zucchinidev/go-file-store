package www

import (
	"github.com/julienschmidt/httprouter"
	"github.com/zucchinidev/go-file-store/infrastructure/engine"
	"net/http"
)

type Conf struct{ Addr string }

func Server(c Conf) *http.Server {
	router := httprouter.New()
	router.GET("/status", status)
	return &http.Server{Addr: c.Addr, Handler: router}
}

type statusResp struct {
	Status string `json:"status"`
}

var status = func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	engine.Respond(w, r, http.StatusOK, statusResp{Status: "UP"})
}
