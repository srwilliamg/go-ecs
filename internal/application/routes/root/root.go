package routerRoot

import (
	"net/http"

	"github.com/go-chi/chi"
)

type Root struct {
	router *chi.Mux
}

func NewRootRouter(router *chi.Mux) *Root {
	rootRouter := &Root{
		router: router,
	}

	rootRouter.initRoutes()
	return rootRouter
}

func (u *Root) initRoutes() {
	u.router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("ok"))

		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	})

}
