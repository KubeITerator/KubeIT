package db

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	v1alpha2 "kubeIT/pkg/proto"
)

func (db *Database) AddUser(user *v1alpha2.User) (id primitive.ObjectID, err error) {

	user.Id = ""

	res, e := db.collections.users.InsertOne(db.ctx, user)

	if e != nil {
		return primitive.ObjectID{}, e
	}

	return res.InsertedID.(primitive.ObjectID), e
}

func (db *Database) RemoveUser(userid primitive.ObjectID) (int64, error) {
	res, err := db.collections.users.DeleteOne(db.ctx, bson.M{"_id": userid})

	if res == nil {
		return 0, err
	} else {
		return res.DeletedCount, err
	}
}

func (db *Database) GetUserByID(id primitive.ObjectID) (user *v1alpha2.User, err error) {
	user = &v1alpha2.User{}
	err = db.collections.users.FindOne(db.ctx, bson.M{"_id": id}).Decode(user)

	return user, err
}

func (db *Database) GetTokensByUser(userid primitive.ObjectID) ([]*v1alpha2.Token, error) {
	user, err := db.GetUserByID(userid)
	if user != nil {
		return user.Tokens, err
	} else {
		return nil, err
	}

}

func (db *Database) AddTokenToUser(token *v1alpha2.Token, id primitive.ObjectID) error {

	token.Id = primitive.NewObjectID().Hex()
	_, err := db.collections.users.UpdateByID(db.ctx, id, bson.M{
		"$push": bson.M{"tokens": token},
	})
	if err != nil {

		user, err := db.GetUserByID(id)

		if err != nil {
			return err
		}

		if user.Tokens == nil {
			_, err = db.collections.users.UpdateByID(db.ctx, id, bson.M{
				"$set": bson.M{"tokens": bson.A{token}},
			})
		}

		return err
	}

	return err
}

func (db *Database) RemoveTokenFromUser(userid, tokenid primitive.ObjectID) error {

	_, err := db.collections.users.UpdateByID(db.ctx, userid, bson.M{
		"$pull": bson.M{"tokens": bson.M{"_id": tokenid.Hex()}},
	})

	return err
}

func (db *Database) AddUserToGroup(userid, groupid primitive.ObjectID, permlevel v1alpha2.GrpPermissionLevel) error {

	perm := v1alpha2.GroupPermission{
		GroupId:    groupid.Hex(),
		Permission: permlevel,
		ProjectIds: []string{},
	}

	userobj, err := db.GetUserByID(userid)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.New("unknown group id")
		}
		return err
	}

	var result bson.M
	err = db.collections.groups.FindOne(db.ctx, bson.D{{"_id", groupid}}).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.New("unknown group id")
		}
		return err
	}

	callback := func(sessCtx mongo.SessionContext) (interface{}, error) {
		// Important: You must pass sessCtx as the Context parameter to the operations for them to be executed in the
		// transaction.
		//userperm := v1alpha2.UserPerms{
		//	UserId:    userid.Hex(),
		//	PermLevel: perm.Permission,
		//}

		userperm := v1alpha2.UserPerms{
			UserId:    userid.Hex(),
			UserName:  userobj.Name,
			PermLevel: perm.Permission,
		}

		_, err := db.collections.groups.UpdateOne(db.ctx, bson.M{"_id": groupid, "userperms.userid": bson.M{"$ne": userid.Hex()}}, bson.M{
			"$addToSet": bson.M{"userperms": &userperm},
		})

		if err != nil {
			return nil, err
		}

		_, err = db.collections.users.UpdateByID(db.ctx, userid, bson.M{
			"$addToSet": bson.M{"gpermissions": &perm},
		})
		if err != nil {

			user, err := db.GetUserByID(userid)

			if err != nil {
				return nil, err
			}

			if user.GPermissions == nil {
				_, err = db.collections.users.UpdateByID(db.ctx, userid, bson.M{
					"$set": bson.M{"gpermissions": bson.A{perm}},
				})
			}

			return nil, err
		}

		return nil, err
	}
	// Step 2: Start a session and run the callback using WithTransaction.
	session, err := db.client.StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(db.ctx)
	_, err = session.WithTransaction(db.ctx, callback)

	return err

}

func (db *Database) RemoveUserFromGroup(userid, groupID primitive.ObjectID) error {

	callback := func(sessCtx mongo.SessionContext) (interface{}, error) {
		// Important: You must pass sessCtx as the Context parameter to the operations for them to be executed in the
		// transaction.

		ret, err := db.collections.groups.UpdateByID(db.ctx, groupID, bson.M{
			"$pull": bson.M{"userperms": bson.M{"userid": userid.Hex()}},
		})

		if ret != nil && ret.ModifiedCount == 0 {
			return nil, errors.New("no items found for update")
		}

		if err != nil {
			return nil, err
		}

		ret, err = db.collections.users.UpdateByID(db.ctx, userid, bson.M{
			"$pull": bson.M{"gpermissions": bson.M{"groupid": groupID.Hex()}},
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

func (db *Database) UpdateUserGrpPermLevel(userid, grpid primitive.ObjectID, perm v1alpha2.GrpPermissionLevel) error {
	callback := func(sessCtx mongo.SessionContext) (interface{}, error) {

		ret, err := db.collections.groups.UpdateOne(db.ctx, bson.M{"_id": grpid, "userperms.userid": userid.Hex()}, bson.M{
			"$set": bson.M{"userperms.$.permlevel": perm},
		})
		if ret != nil && ret.ModifiedCount == 0 {
			return nil, errors.New("no items found for update")
		}

		if err != nil {
			return nil, err
		}

		ret, err = db.collections.users.UpdateOne(db.ctx, bson.M{"_id": userid, "gpermissions.groupid": grpid.Hex()}, bson.M{
			"$set": bson.M{"gpermissions.$.permission": &perm},
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

func (db *Database) GetGroupPermissionFromUser(userid, groupid primitive.ObjectID) (*v1alpha2.GroupPermission, error) {

	user, err := db.GetUserByID(userid)

	if user == nil {
		return nil, errors.New("no user found")
	}
	for _, perm := range user.GPermissions {
		if perm.GroupId == groupid.Hex() {
			return perm, nil
		}
	}

	return nil, err
}

func (db *Database) IsUserInGroup(uid, grpid primitive.ObjectID) (bool, error) {
	result := db.collections.groups.FindOne(db.ctx, bson.D{{"_id", grpid}, {"userperms.userid", uid.Hex()}})

	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, result.Err()
	}

	return true, nil
}

func (db *Database) AddUserToProject(uid, pid, grpid primitive.ObjectID) error {

	isInGrp, err := db.IsUserInGroup(uid, grpid)
	if err != nil || isInGrp == false {
		return errors.New("user not in group")
	}

	callback := func(sessCtx mongo.SessionContext) (interface{}, error) {

		ret, err := db.collections.groups.UpdateOne(db.ctx, bson.D{{"projects._id", pid.Hex()}}, bson.M{
			"$addToSet": bson.M{"projects.$.userids": uid.Hex()},
		})

		if ret != nil && ret.ModifiedCount == 0 {
			return nil, errors.New("no items found for update")
		}

		if err != nil {
			return nil, err
		}

		ret, err = db.collections.users.UpdateByID(db.ctx, bson.M{"_id": uid, "gpermissions.projectids": pid.Hex()}, bson.M{
			"$addToSet": bson.M{"gpermissions.$.projectids": pid.Hex()},
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

func (db *Database) RemoveUserFromProject(uid, pid primitive.ObjectID) error {

	callback := func(sessCtx mongo.SessionContext) (interface{}, error) {

		ret, err := db.collections.groups.UpdateOne(db.ctx, bson.D{{"projects._id", pid.Hex()}}, bson.M{
			"$pull": bson.M{"projects.$.userids": uid.Hex()},
		})

		if ret != nil && ret.ModifiedCount == 0 {
			return nil, errors.New("no items found for update")
		}

		if err != nil {
			return nil, err
		}

		ret, err = db.collections.users.UpdateByID(db.ctx, bson.M{"_id": uid, "gpermissions.projectids": pid.Hex()}, bson.M{
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
