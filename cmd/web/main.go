package main

import (
	"database/sql"
	"flag"
	"food_market/pkg/models/dbs"
	"github.com/golangcollege/sessions"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	session       *sessions.Session
	templateCache map[string]*template.Template
	user          *dbs.UserModel
	product       *dbs.ProductModel
	category      *dbs.CategoryModel
	//history       *dbs.OrderHistoryModel
}

func main() {
	//dsn := "user=food_market_d8v1_user password=0p5Y4mvLHVfJLUCJAGHeUSaa8sAIp5aL dbname=food_market_d8v1 sslmode=disable host=dpg-cnhkrhed3nmc739f2r30-a port=5432"
	dsn := "user=bxit password=aa dbname=bxit sslmode=disable host=localhost port=5433"
	addr := flag.String("addr", ":4001", "HTTP network address")

	secret := flag.String("secret", "s6Ndh+pPbnzHbS*+9Pk8qGWhTzbpa@ge", "Secret key")

	flag.Parse()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowCredentials: true,
	})

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	templateCache, err := newTemplateCache("./ui/html/")
	if err != nil {
		errorLog.Fatal(err)
	}

	session := sessions.New([]byte(*secret))
	session.Lifetime = 12 * time.Hour

	app := &application{
		errorLog:      errorLog,
		infoLog:       infoLog,
		session:       session,
		templateCache: templateCache,
		user:          &dbs.UserModel{DB: db},
		product:       &dbs.ProductModel{DB: db},
		category:      &dbs.CategoryModel{DB: db},
		//history:       &dbs.OrderHistoryModel{DB: db},
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  c.Handler(app.routes()),
		// TLSConfig:    tlsConfig,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	infoLog.Printf("Starting server on %s", *addr)

	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
