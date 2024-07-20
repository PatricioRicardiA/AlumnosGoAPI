package main

import (
	"GoApiEchoTest3/db"
	"GoApiEchoTest3/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	db.ConexionDB()

	AlumnosCollection := db.Client.Database("Escuelita").Collection("Alumnos")
	handlers.SetAlumnosCollection(AlumnosCollection)

	e.POST("/Alumnos", handlers.CrearAlumno)
	e.GET("/Alumnos", handlers.ObtenerTodosAlumnos)
	e.GET("/Alumnos/:id", handlers.ObtenerAlumnoxId)
	e.PUT("/Alumnos/:id", handlers.ActualizarAlumno)
	e.DELETE("/Alumnos/:id", handlers.EliminarAlumno)

	e.Start(":8080")
}
