package main

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func getChanges(c *gin.Context) {
	cur, err := db.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer cur.Close(ctx)
	changes := make([]Change, 0)

	for cur.Next(ctx) {
		var change Change
		cur.Decode(&change)
		changes = append(changes, change)
	}

	c.JSON(http.StatusOK, changes)
}

func postChanges(c *gin.Context) {
	var change Change
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := c.ShouldBindJSON(&change); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// if time is empty set to now
	if change.Time.IsZero() {
		change.Time = time.Now()
	}

	change.ID = primitive.NewObjectID()
	db.InsertOne(ctx, &change)

	c.JSON(http.StatusOK, gin.H{
		"message": "change created",
		"change":  change,
	})
}
