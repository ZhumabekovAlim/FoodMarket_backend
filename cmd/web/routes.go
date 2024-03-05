package main

import (
	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
	"net/http"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders, makeResponseJSON)

	dynamicMiddleware := alice.New(app.session.Enable)

	mux := pat.New()

	// HOME PAGE
	mux.Get("/", dynamicMiddleware.ThenFunc(app.home))

	// AUTH
	mux.Post("/signup", dynamicMiddleware.ThenFunc(app.signupUser))
	mux.Post("/login", dynamicMiddleware.ThenFunc(app.loginUser))
	mux.Post("/logout", dynamicMiddleware.ThenFunc(app.logOut))

	// USER
	mux.Patch("/admin", dynamicMiddleware.ThenFunc(app.updateUser))

	// ADMIN
	mux.Get("/admin/users", dynamicMiddleware.ThenFunc(app.getAllUsers))

	// USER PROFILE
	mux.Get("/profile/:id", dynamicMiddleware.ThenFunc(app.profile))

	// PRODUCT
	mux.Post("/create-product", dynamicMiddleware.ThenFunc(app.createProduct))
	mux.Get("/products", dynamicMiddleware.ThenFunc(app.products))
	mux.Get("/product/:id", dynamicMiddleware.ThenFunc(app.productsById))
	mux.Patch("/update-product", dynamicMiddleware.ThenFunc(app.updateProduct))
	mux.Del("/delete-product", dynamicMiddleware.ThenFunc(app.deleteProduct))

	// CATEGORY
	mux.Post("/create-category", dynamicMiddleware.ThenFunc(app.createCategory))
	mux.Get("/categories", dynamicMiddleware.ThenFunc(app.categories))
	mux.Get("/category/:id", dynamicMiddleware.ThenFunc(app.categoryById))
	mux.Patch("/update-category", dynamicMiddleware.ThenFunc(app.updateCategory))
	mux.Del("/delete-category", dynamicMiddleware.ThenFunc(app.deleteCategory))

	//mux.Get("/histories", dynamicMiddleware.ThenFunc(app.histories))
	//mux.Get("/history/:id", dynamicMiddleware.ThenFunc(app.historyById))
	//mux.Post("/histories-user", dynamicMiddleware.ThenFunc(app.historiesByUserId))
	//mux.Post("/create-history", dynamicMiddleware.ThenFunc(app.createHistory))
	//mux.Patch("/update-history", dynamicMiddleware.ThenFunc(app.updateHistory))
	//mux.Del("/delete-history", dynamicMiddleware.ThenFunc(app.deleteHistory))

	return standardMiddleware.Then(mux)
}
