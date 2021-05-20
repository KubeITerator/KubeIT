package db

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	v1alpha2 "kubeIT/pkg/proto"
)

func (db *Database) AddTask(task *v1alpha2.Task) (id primitive.ObjectID, err error) {

	task.Id = ""

	res, e := db.collections.tasks.InsertOne(db.ctx, task)

	if e != nil {
		return primitive.ObjectID{}, e
	}

	return res.InsertedID.(primitive.ObjectID), e
}

func (db *Database) TaskExists(taskid primitive.ObjectID) bool {
	num, _ := db.collections.tasks.CountDocuments(db.ctx, bson.D{{"_id", taskid}})

	if num > 0 {
		return true
	} else {
		return false
	}
}

func (db *Database) GetTask(taskid primitive.ObjectID) (t *v1alpha2.Task, err error) {
	task := &v1alpha2.Task{}
	err = db.collections.tasks.FindOne(db.ctx, bson.M{"_id": taskid}).Decode(task)
	return task, err
}

func (db *Database) DeleteTask(taskid primitive.ObjectID) (int64, error) {
	res, err := db.collections.tasks.DeleteOne(db.ctx, bson.M{"_id": taskid})

	if res == nil {
		return 0, err
	} else {
		return res.DeletedCount, err
	}
}
