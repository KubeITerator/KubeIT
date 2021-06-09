package db

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"kubeIT/pkg/grpc/storage"
)

func (db *Database) AddStorage(store *storage.File) (id primitive.ObjectID, err error) {

	store.Id = ""

	res, e := db.collections.storage.InsertOne(db.ctx, store)

	if e != nil {
		return primitive.ObjectID{}, e
	}

	return res.InsertedID.(primitive.ObjectID), e
}

func (db *Database) StorageExist(storeid primitive.ObjectID) bool {
	num, _ := db.collections.storage.CountDocuments(db.ctx, bson.D{{"_id", storeid}})

	if num > 0 {
		return true
	} else {
		return false
	}
}

func (db *Database) GetFile(storeid primitive.ObjectID) (w *storage.File, err error) {
	f := &storage.File{}
	err = db.collections.storage.FindOne(db.ctx, bson.M{"_id": storeid}).Decode(f)
	return f, err
}

func (db *Database) DeleteFile(fileid primitive.ObjectID) (int64, error) {
	res, err := db.collections.storage.DeleteOne(db.ctx, bson.M{"_id": fileid})

	if res == nil {
		return 0, err
	} else {
		return res.DeletedCount, err
	}
}
