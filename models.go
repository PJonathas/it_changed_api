package main

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Change struct {
	ID          primitive.ObjectID `json:"id"`
	Description string             `json:"description" binding:"required"`
	Type        string             `json:"type" binding:"required"`
	Time        time.Time          `json:"time"`
}
