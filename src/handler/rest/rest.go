package rest

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"github.com/gorilla/mux"
	"github.com/trongtb88/wagerservice/src/business/usecase"
)

// REST rest interface
type REST interface{}

var once = &sync.Once{}

type rest struct {
	logger log.Logger
	mux    *mux.Router
	uc     *usecase.Usecase
}

func Init(logger log.Logger,  router *mux.Router, uc *usecase.Usecase) REST {
	var e *rest
	once.Do(func() {
		e = &rest{
			logger: logger,
			mux:    router,
			uc:     uc,
		}
		e.Serve()
	})
	return e
}


func (rst *rest) Serve() {

	rst.mux.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	})

	rst.mux.HandleFunc("/wagers", rst.CreateWager).Methods(http.MethodPost)
	rst.mux.HandleFunc("/buy/{wager_id}", rst.BuyWager).Methods(http.MethodPost)
	rst.mux.HandleFunc("/wagers", rst.GetWagers).Queries("limit", "{param2:[0-9,]+}", "page", "{param1:[0-9,]+}").Methods(http.MethodGet)
}

func (rst *rest) httpRespSuccess(w http.ResponseWriter, statusCode int, resp interface{}) {
	raw, _ := json.Marshal( resp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, _ = w.Write(raw)
}

func (rst *rest) httpRespError(w http.ResponseWriter, statusCode int, err error, message string) {
	var e ErrorMessage
	if err != nil {
		e = ErrorMessage{
			Message: message,
		}
	} else {
		e = ErrorMessage{
			Message: message,
		}
	}

	raw, _ := json.Marshal(e)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, _ = w.Write(raw)
}