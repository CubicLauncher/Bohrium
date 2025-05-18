package main

import (
	"fmt"
	"net/http"
	"os"

	"Bohrium/handlers"
	"Bohrium/services"
)

func main() {
	// Inicializar servicios
	minecraftService := services.NewMinecraftService()

	// Inicializar handlers
	skinHandler := handlers.NewSkinHandler(minecraftService)

	// Configurar rutas
	http.HandleFunc("/skin/", skinHandler.HandleSkinRequest)

	// Configurar puerto
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Iniciar servidor
	fmt.Printf("Servidor iniciado en http://localhost:%s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Printf("Error al iniciar el servidor: %v\n", err)
		os.Exit(1)
	}
} 
