package handlers

import (
	"context"
	"net/http"
	"time"

	"GoApiEchoTest3/models"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var AlumnosCollection *mongo.Collection

func SetAlumnosCollection(c *mongo.Collection) {
	AlumnosCollection = c
}

func CrearAlumno(c echo.Context) error {
	var Alumno models.Alumno
	if err := c.Bind(&Alumno); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := AlumnosCollection.InsertOne(ctx, Alumno)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, result)
}

func ObtenerTodosAlumnos(c echo.Context) error {

	var Alumnos []models.Alumno

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := AlumnosCollection.Find(ctx, bson.M{})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var Alumno models.Alumno
		cursor.Decode(&Alumno)
		Alumnos = append(Alumnos, Alumno)
	}

	if err := cursor.Err(); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, Alumnos)
}

func ObtenerAlumnoxId(c echo.Context) error {
	idParam := c.Param("id")
	AlumnoId, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var Alumno models.Alumno

	err = AlumnosCollection.FindOne(ctx, bson.M{"_id": AlumnoId}).Decode(&Alumno)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.JSON(http.StatusNotFound, "Alumno no encontrado")
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, Alumno)

}

func ActualizarAlumno(c echo.Context) error {
	idParam := c.Param("id")
	AlumnoID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	var Alumno models.Alumno
	if err := c.Bind(&Alumno); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	update := bson.M{
		"$set": Alumno,
	}

	_, err = AlumnosCollection.UpdateOne(ctx, bson.M{"_id": AlumnoID}, update)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "El alumno fue actualizao")
}

func EliminarAlumno(c echo.Context) error {
	idParam := c.Param("id")
	AlumnoID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = AlumnosCollection.DeleteOne(ctx, bson.M{"_id": AlumnoID})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "Alumno eliminado")

}
