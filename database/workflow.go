package db

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"kubeIT/pkg/grpc/workflow"
)

func (db *Database) AddWorkflow(wf *workflow.Workflow) (id primitive.ObjectID, err error) {

	wf.Id = ""

	res, e := db.collections.workflows.InsertOne(db.ctx, wf)

	if e != nil {
		return primitive.ObjectID{}, e
	}

	return res.InsertedID.(primitive.ObjectID), e
}

func (db *Database) WorkflowExists(wfid primitive.ObjectID) bool {
	num, _ := db.collections.workflows.CountDocuments(db.ctx, bson.D{{"_id", wfid}})

	if num > 0 {
		return true
	} else {
		return false
	}
}

func (db *Database) GetWorkflow(wfid primitive.ObjectID) (w *workflow.Workflow, err error) {
	wf := &workflow.Workflow{}
	err = db.collections.workflows.FindOne(db.ctx, bson.M{"_id": wfid}).Decode(wf)
	return wf, err
}

func (db *Database) DeleteWorkflow(wfid primitive.ObjectID) (int64, error) {
	res, err := db.collections.workflows.DeleteOne(db.ctx, bson.M{"_id": wfid})

	if res == nil {
		return 0, err
	} else {
		return res.DeletedCount, err
	}
}
