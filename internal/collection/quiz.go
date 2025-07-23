package collection

import (
	"context"

	"github.com/JagdeepSingh13/go_quiz/internal/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type QuizCollection struct {
	collection *mongo.Collection
}

func Quiz(collection *mongo.Collection) *QuizCollection {
	return &QuizCollection{
		collection: collection,
	}
}

func (c QuizCollection) InsertQuiz(quiz entity.Quiz) error {
	_, err := c.collection.InsertOne(context.Background(), quiz)
	return err
}

func (c QuizCollection) GetQuizById(id primitive.ObjectID) (*entity.Quiz, error) {
	result := c.collection.FindOne(context.Background(), bson.M{"_id": id})

	var quiz entity.Quiz
	err := result.Decode(&quiz)
	if err != nil {
		return nil, err
	}

	return &quiz, nil
}

func (c QuizCollection) GetQuizzes() ([]entity.Quiz, error) {
	cur, err := c.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}

	var quiz []entity.Quiz
	err = cur.All(context.Background(), &quiz)
	if err != nil {
		return nil, err
	}

	return quiz, nil
}
