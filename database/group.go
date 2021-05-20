package db

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"kubeIT/pkg/grpc/group"
)

func (db *Database) AddGroup(grp *group.Group) (id primitive.ObjectID, err error) {
	grp.GroupId = ""
	res, e := db.collections.groups.InsertOne(db.ctx, grp)

	if e != nil {
		return primitive.ObjectID{}, e
	}

	return res.InsertedID.(primitive.ObjectID), e
}

func (db *Database) RemoveGroup(grpid primitive.ObjectID) error {
	_, err := db.collections.groups.DeleteOne(db.ctx, bson.M{"_id": grpid})
	if err != nil {
		return err
	}
	_, err = db.collections.users.UpdateMany(db.ctx, bson.M{}, bson.M{
		"$pull": bson.M{"gpermissions": bson.M{"groupid": grpid.Hex()}}})

	return err
}

func (db *Database) GetGroupByID(grpid primitive.ObjectID) (grp *group.Group, err error) {
	grp = &group.Group{}
	err = db.collections.groups.FindOne(db.ctx, bson.M{"_id": grpid}).Decode(grp)

	return grp, err
}

func (db *Database) AddProject(project *group.Project, groupid, userid primitive.ObjectID) (projectid *primitive.ObjectID, e error) {
	newpid := primitive.NewObjectID()

	project.ProjectId = newpid.Hex()
	callback := func(sessCtx mongo.SessionContext) (interface{}, error) {

		ret, err := db.collections.groups.UpdateByID(db.ctx, groupid, bson.M{
			"$push": bson.M{"projects": project},
		})

		if ret != nil && ret.ModifiedCount == 0 {
			return nil, errors.New("no items found for update")
		}

		if err != nil {
			return nil, err
		}

		ret, err = db.collections.users.UpdateOne(db.ctx, bson.M{"_id": userid, "gpermissions.groupid": groupid.Hex()}, bson.M{
			"$push": bson.M{"gpermissions.$.projectids": newpid.Hex()},
		})

		if ret != nil && ret.ModifiedCount == 0 {
			return nil, errors.New("no items found for update")
		}

		return nil, err
	}
	session, err := db.client.StartSession()
	if err != nil {
		return nil, err
	}
	defer session.EndSession(db.ctx)
	_, err = session.WithTransaction(db.ctx, callback)

	return &newpid, err

}

func (db *Database) GetProjectByID(pid primitive.ObjectID) (proj *group.Project, err error) {
	grp := &group.Group{}

	err = db.collections.groups.FindOne(db.ctx, bson.D{{"projects._id", pid.Hex()}}).Decode(grp)

	if err != nil {
		return nil, err
	}

	for _, project := range grp.Projects {
		if project.ProjectId == pid.Hex() {
			return project, nil
		}
	}
	return nil, nil
}

func (db *Database) DeleteProject(pid primitive.ObjectID) error {
	callback := func(sessCtx mongo.SessionContext) (interface{}, error) {

		ret, err := db.collections.groups.UpdateOne(db.ctx, bson.D{{"projects._id", pid.Hex()}}, bson.M{
			"$pull": bson.M{"projects": bson.M{"_id": pid.Hex()}},
		})

		if ret != nil && ret.ModifiedCount == 0 {
			return nil, errors.New("no items found for update")
		}

		if err != nil {
			return nil, err
		}

		ret, err = db.collections.users.UpdateMany(db.ctx, bson.M{"gpermissions.projectids": pid.Hex()}, bson.M{
			"$pull": bson.M{"gpermissions.$.projectids": pid.Hex()},
		})
		if ret != nil && ret.ModifiedCount == 0 {
			return nil, errors.New("no items found for update")
		}

		return nil, err
	}
	session, err := db.client.StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(db.ctx)
	_, err = session.WithTransaction(db.ctx, callback)

	return err

}
