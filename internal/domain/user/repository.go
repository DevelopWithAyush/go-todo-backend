package user

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/developwithayush/go-todo-app/internal/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository interface {
	FindByEmail(ctx context.Context, email string) (*User, error)
	UpsertOTP(ctx context.Context, email string, otpHash string, expiresAt time.Time) (*User, error)
	ClearOTP(ctx context.Context, userID primitive.ObjectID) error
}

type repo struct{}

func NewRepository() Repository {
	return &repo{}
}

func (r *repo) FindByEmail(ctx context.Context, email string) (*User, error) {
	var user User
	err := db.Users.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}

	// Print user data in JSON format (similar to Express console.log)
	userJSON, _ := json.MarshalIndent(user, "", "  ")
	fmt.Println("User:", string(userJSON))

	return &user, nil
}

func (r *repo) UpsertOTP(ctx context.Context, email, otpHash string, expiresAt time.Time) (*User, error) {
	now := time.Now()
	update := bson.M{
		"$set": bson.M{
			"email":        email,
			"otpHash":      otpHash,
			"otpExpiresAt": expiresAt,
			"updatedAt":    now,
		},
		"$setOnInsert": bson.M{
			"createdAt": now,
		},
	}

	opt := options.FindOneAndUpdate().
		SetUpsert(true).
		SetReturnDocument(options.After)

	var user User
	err := db.Users.FindOneAndUpdate(ctx, bson.M{"email": email}, update, opt).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repo) ClearOTP(ctx context.Context, userID primitive.ObjectID) error {
	_, err := db.Users.UpdateOne(ctx, userID, bson.M{
		"$set": bson.M{
			"otpHash":      "",
			"otpExpiresAt": time.Time{},
		},
	})

	return err
}
