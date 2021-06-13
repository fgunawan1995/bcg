package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/rs/cors"

	"github.com/fgunawan1995/bcg/common"
	"github.com/fgunawan1995/bcg/config"
	cachedal "github.com/fgunawan1995/bcg/dal/cache"
	dbdal "github.com/fgunawan1995/bcg/dal/db"
	"github.com/fgunawan1995/bcg/resolver"
	"github.com/fgunawan1995/bcg/resources"
	"github.com/fgunawan1995/bcg/schema"
	"github.com/fgunawan1995/bcg/usecase"
)

func main() {
	// Init DB
	db, err := resources.ConnectDb(config.GetConfig())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Init cache
	cache := resources.InitCache()

	// Init schema
	dbDAL := dbdal.New(db)
	cacheDAL := cachedal.New(cache)
	commonLayer := common.New(cacheDAL, dbDAL)
	usecaseLayer := usecase.New(commonLayer, cacheDAL, dbDAL)
	handlerLayer := resolver.New(usecaseLayer, commonLayer, cacheDAL, dbDAL)
	schema := graphql.MustParseSchema(schema.Schema, handlerLayer)

	// Init routes
	r := mux.NewRouter()
	r.HandleFunc("/", index).Methods("GET")
	r.Handle("/api/v1/query", &relay.Handler{Schema: schema})

	// Start server
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Content-Type"},
	})
	handler := c.Handler(r)
	port := fmt.Sprintf(":%s", config.GetConfig().Server.Port)
	fmt.Printf("Server started at %s", port)
	log.Fatal(http.ListenAndServe(port, handler))
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello"))
}
