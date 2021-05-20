package db

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	v1alpha2 "kubeIT/pkg/proto"
)

func (db *Database) AddSchema(schema *v1alpha2.Schema) (id primitive.ObjectID, err error) {

	schema.Id = ""

	res, e := db.collections.schema.InsertOne(db.ctx, schema)

	if e != nil {
		return primitive.ObjectID{}, e
	}

	return res.InsertedID.(primitive.ObjectID), e
}

func (db *Database) SchemaExists(schemaid primitive.ObjectID) bool {
	num, _ := db.collections.schema.CountDocuments(db.ctx, bson.D{{"_id", schemaid}})

	if num > 0 {
		return true
	} else {
		return false
	}
}

func (db *Database) GetSchema(schemaid primitive.ObjectID) (s *v1alpha2.Schema, err error) {
	schema := &v1alpha2.Schema{}
	err = db.collections.schema.FindOne(db.ctx, bson.M{"_id": schemaid}).Decode(schema)
	return schema, err
}

func (db *Database) DeleteSchema(schemaid primitive.ObjectID) (int64, error) {
	res, err := db.collections.schema.DeleteOne(db.ctx, bson.M{"_id": schemaid})

	if res == nil {
		return 0, err
	} else {
		return res.DeletedCount, err
	}
}
