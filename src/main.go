package main

import (
	"aurora-insights-api/src/controllers"
	"aurora-insights-api/src/models"
	"aurora-insights-api/src/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	controllers.RegisterControllers(r)
	fmt.Println("Server listening on port 8000...")
	http.ListenAndServe(":8000", r)
}
