package Models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type GroupMessage struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	MessageFrom primitive.ObjectID `bson:"messageFrom"`
	Content     string             `bson:"content"`
	CreatedBy   primitive.ObjectID `bson:"createdBy"`
	CreatedAt   time.Time          `bson:"createdAt"`
	Liked       uint               `bson:"liked"`
	Disliked    uint               `bson:"disliked"`
	Deleted     bool               `bson:"deleted"`
}
