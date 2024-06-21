package Models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type MessageBoard struct {
	ID          primitive.ObjectID ` bson:"_id,omitempty"`
	MessageFrom string             `bson:"messageFrom"`
	Content     string             `bson:"content"`
	CreatedAt   time.Time          `bson:"createdAt"`
	Liked       *uint              `bson:"liked,omitempty"`
	Disliked    *uint              `bson:"disliked,omitempty"`
	Deleted     *bool              `bson:"deleted,omitempty"`
}
