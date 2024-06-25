package Models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Message struct {
	ID          primitive.ObjectID ` bson:"_id,omitempty"`
	MessageFrom string             `bson:"messageFrom"`
	MessageTo   string             `bson:"messageTo"`
	Content     string             `bson:"content"`
	CreatedAt   time.Time          `bson:"createdAt"`
	Seen        *bool              `bson:"seen"`
	Liked       *uint              `bson:"liked"`
	Disliked    *uint              `bson:"disliked"`
	Deleted     *bool              `bson:"deleted"`
}
