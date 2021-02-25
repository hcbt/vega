package db

import (
	//"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/bson/primitive"
)

type Track struct {
	ID      string   `bson:"_id"`
	Artist  string   `bson:"artist,omitempty"`
	Title   string   `bson:"title,omitempty"`
	Genre   string   `bson:"genre,omitempty"`
	Tags    []string `bson:"tags,omitempty"`
}
