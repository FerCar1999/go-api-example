package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	portstring := os.Getenv("APP_PORT")
	if portstring == "" {
		log.Fatal("No se ha encontrado el puerto en las variables de entorno")
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1 := chi.NewRouter()

	v1.Get("/healthz", handlerReadiness)
	v1.Get("/healthz2", handlerError)

	router.Mount("/v1", v1)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + portstring,
	}

	log.Printf("Server corriendo en http://localhost:%v", portstring)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
