package main

import (
	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
	"net/http"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	dynamicMiddleware := alice.New(app.session.Enable)

	mux := pat.New()

	mux.Get("/", dynamicMiddleware.ThenFunc(app.home))
	mux.Get("/test", dynamicMiddleware.ThenFunc(app.products))
	mux.Get("/profile", dynamicMiddleware.ThenFunc(app.profile))
	mux.Get("/categories", dynamicMiddleware.ThenFunc(app.categories))
	mux.Post("/create_category", dynamicMiddleware.ThenFunc(app.createCategory))
	mux.Post("/create_product", dynamicMiddleware.ThenFunc(app.createProduct))
	mux.Patch("/update_product", dynamicMiddleware.ThenFunc(app.updateProduct))
	mux.Del("/delete_product", dynamicMiddleware.ThenFunc(app.deleteProduct))
	mux.Patch("/update_category", dynamicMiddleware.ThenFunc(app.updateCategory))
	mux.Del("/delete_category", dynamicMiddleware.ThenFunc(app.deleteCategory))

	return standardMiddleware.Then(mux)
}
