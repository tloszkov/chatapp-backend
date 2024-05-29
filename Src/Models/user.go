package Models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func main() {
	type Message struct {
		ID        primitive.ObjectID ` bson:"_id,omitempty"`
		FirstName string             `bson:"firstName"`
		LastName  string             `bson:"lastName"`
		Email     string             `bson:"email"`
		Password  string             `bson:"password"`
		CreatedAt time.Time          `bson:"createdAt"`
		Deleted   bool               `bson:"deleted"`
	}

}
