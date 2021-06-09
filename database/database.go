package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Collections struct {
	users, groups, schema, tasks, workflows, storage *mongo.Collection
}

type Database struct {
	client      *mongo.Client
	ctx         context.Context
	database    *mongo.Database
	collections Collections
}

func (db *Database) Init(credential options.Credential, uri string) error {

	db.ctx = context.Background()
	var err error
	db.client, err = mongo.Connect(db.ctx, options.Client().ApplyURI(uri).SetAuth(credential))
	if err != nil {
		return err
	}
	db.database = db.client.Database("testdb")
	// initialize collections -> Access methods for each collection is stored in separate go-file ({collection-name}.go)
	db.collections.users = db.database.Collection("users")
	db.collections.groups = db.database.Collection("groups")
	db.collections.schema = db.database.Collection("schema")
	db.collections.tasks = db.database.Collection("tasks")
	db.collections.workflows = db.database.Collection("workflows")
	db.collections.storage = db.database.Collection("storage")
	return err
}

func (db *Database) Stop() error {
	err := db.client.Disconnect(db.ctx)
	return err
}
