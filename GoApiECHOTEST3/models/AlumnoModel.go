package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Alumno struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Nombre    string             `json:"Nombre" bson:"Nombre"`
	Matricula string             `json:"Matricula" bson:"Matricula"`
	Grado     int                `json:"Grado" bson:"Grado"`
}
