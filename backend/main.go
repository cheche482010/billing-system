package main

import (
	"fmt"
	"log"
	"net/http"

	"billing-system/db"
	"billing-system/routes"

	color "github.com/fatih/color"
)

func main() {

	if !db.Connect() {
		color.Red("Error en Compilación")
		return
	} else {
		color.Green("Compilación Exitosa!")
	}

	fmt.Println("Ahora puede ver el Backend en el navegador.")
	color.Set(color.FgWhite)
	fmt.Print("Local: ")
	color.Set(color.FgBlue)
	fmt.Println("http://localhost:8080")
	color.Unset()

	err := routes.InitializeRoutes()
	if err != nil {
		log.Fatalf("Error initializing routes: %v", err)
	}

	log.Fatal(http.ListenAndServe(":8080", routes.Router))
}
