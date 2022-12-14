package controller

import (
	"net/http"
	"net/http/httptest"

	es "github.com/calogxro/qaservice/db/event_store"
	rr "github.com/calogxro/qaservice/db/read_repository"
	"github.com/calogxro/qaservice/domain"
	"github.com/calogxro/qaservice/service"
	"github.com/gin-gonic/gin"
)

var testAnswer = domain.Answer{Key: "name", Value: "John"}
var testJson = []byte(`{"key": "name", "value": "John"}`)

type TestRouter struct {
	router *gin.Engine
}

func NewTestRouter() *TestRouter {
	es := es.NewEventStoreStub()
	rr := rr.NewReadRepositoryStub()
	qaservice := service.NewQAService(es)
	projection := service.NewQAProjection(rr)
	ctrl := NewController(qaservice, projection)
	router := NewRouter(ctrl)

	projector := service.NewProjector(rr)
	es.Subscribe(func(event *domain.Event) {
		projector.Project(event)
	})

	return &TestRouter{
		router: router,
	}
}

func (r *TestRouter) executeReq(req *http.Request) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	r.router.ServeHTTP(w, req)
	return w
}
