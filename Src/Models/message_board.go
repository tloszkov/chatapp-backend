package Models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func MessageBoardModel() {
	type MessageBoard struct {
		ID          primitive.ObjectID ` bson:"_id,omitempty"`
		MessageFrom string             `bson:"messageFrom"`
		Content     string             `bson:"content"`
		CreatedAt   time.Time          `bson:"createdAt"`
		Liked       uint               `bson:"liked"`
		Disliked    uint               `bson:"disliked"`
		Deleted     bool               `bson:"deleted"`
	}

}
