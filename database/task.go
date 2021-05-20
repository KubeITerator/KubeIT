package db

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"kubeIT/pkg/grpc/task"
)

func (db *Database) AddTask(tsk *task.Task) (id primitive.ObjectID, err error) {

	tsk.Id = ""

	res, e := db.collections.tasks.InsertOne(db.ctx, tsk)

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

func (db *Database) GetTask(taskid primitive.ObjectID) (t *task.Task, err error) {
	tsk := &task.Task{}
	err = db.collections.tasks.FindOne(db.ctx, bson.M{"_id": taskid}).Decode(tsk)
	return tsk, err
}

func (db *Database) DeleteTask(taskid primitive.ObjectID) (int64, error) {
	res, err := db.collections.tasks.DeleteOne(db.ctx, bson.M{"_id": taskid})

	if res == nil {
		return 0, err
	} else {
		return res.DeletedCount, err
	}
}
