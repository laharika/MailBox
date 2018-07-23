package main

import(
    
    "mailbox_server/handlers"
    "github.com/julienschmidt/httprouter"
    "github.com/rs/cors"
    "net/http"
    "log"
)
func main() {
	router := httprouter.New()
	handlers.ConfigureRoutes(router)

	// Add CORS support (CrossError Origin Resource Sharing)
    c := cors.New(cors.Options{
    		AllowedOrigins: []string{"*"},
    		AllowedMethods: []string{"GET", "POST", "DELETE"},
    	})
    handler := c.Handler(router)
    log.Fatal(http.ListenAndServe(":8080", handler))
}
