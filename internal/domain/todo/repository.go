package todo

import (
	"context"

	"github.com/developwithayush/go-todo-app/internal/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository interface {
	ListByUser(ctx context.Context, userID primitive.ObjectID) ([]Todo, error)
	Create(ctx context.Context, todo Todo) (*Todo, error)
	Update(ctx context.Context, userID, todoID primitive.ObjectID, update bson.M) error
	Delete(ctx context.Context, userID, todoID primitive.ObjectID) error
	UpdatePosition(ctx context.Context, userID primitive.ObjectID, position int) error
}

type repo struct{}

func NewRepository() Repository {
	return &repo{}
}

func (r *repo) ListByUser(ctx context.Context, userID primitive.ObjectID) ([]Todo, error) {
	opt := options.Find().SetSort(bson.M{"position": 1})

	cur, err := db.Todos.Find(ctx, bson.M{"userId": userID}, opt)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	todos :=[]Todo{} // Initialize as empty slice, not nil
	if err := cur.All(ctx, &todos); err != nil {
		return nil, err
	}
	return todos, nil
}

func (r *repo) Create(ctx context.Context, todo Todo) (*Todo, error) {
	_, err := db.Todos.InsertOne(ctx, todo)
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r *repo) Update(ctx context.Context, userID, todoID primitive.ObjectID, update bson.M) error {
	_, err := db.Todos.UpdateOne(ctx,
		bson.M{"_id": todoID, "userId": userID},
		bson.M{"$set": update})
	if err != nil {
		return err
	}
	return nil
}

func (r *repo) Delete(ctx context.Context, userID, todoID primitive.ObjectID) error {
	_, err := db.Todos.DeleteOne(ctx, bson.M{"_id": todoID, "userId": userID})
	return err
}

func (r *repo) UpdatePosition(ctx context.Context, userID primitive.ObjectID, position int) error {
	_, err := db.Todos.UpdateOne(ctx, bson.M{"userId": userID}, bson.M{"$set": bson.M{"position": position}})
	return err
}
