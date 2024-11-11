package notification

import (
	"context"
	"log"
	"net/http"
	"time"
	"encoding/json"


	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Notification struct {
	ID        string    `bson:"_id,omitempty" json:"id"`
	UserID	  string    `bson:"user_id" json:"user_id"`
	Title     string    `bson:"title" json:"title"`
	Body      string    `bson:"body" json:"body"`
	Read      bool      `bson:"read" json:"read"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
}

var mongoClient *mongo.Client
var notificationsCollection *mongo.Collection

func InitMongoDB(uri string) error {
	var err error
	mongoClient, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}
	notificationsCollection = mongoClient.Database("notificationDB").Collection("notifications")
	return nil
}

func HandleMessage(userID string, messageBody string) error {

	notification := Notification{
		Title:     "Expire Soon",
		UserID:   userID,
		Body:      messageBody,
		Read:      false,
		CreatedAt: time.Now(),
	}

	_, err := notificationsCollection.InsertOne(context.TODO(), notification)
	if err != nil {
		log.Printf("Error saving notification to MongoDB: %v", err)
		return err
	}

	log.Printf("Push notification sent and stored successfully!")
	return nil
}

func GetNotifications(w http.ResponseWriter, r *http.Request) {
	var notifications []Notification
	cursor, err := notificationsCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		http.Error(w, "Failed to retrieve notifications", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var notification Notification
		if err := cursor.Decode(&notification); err != nil {
			log.Printf("Error decoding notification: %v", err)
			continue
		}
		notifications = append(notifications, notification)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(notifications); err != nil {
		http.Error(w, "Failed to encode notifications", http.StatusInternalServerError)
	}
}

func MarkNotificationAsRead(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing notification ID", http.StatusBadRequest)
		return
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"read": true}}

	_, err := notificationsCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		http.Error(w, "Failed to mark notification as read", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Notification marked as read"}`))
}
