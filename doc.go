// Package mr eases interactions with MongoDB.
//
// This package is an opinionated abstraction for "mgo". Documentation for mgo
// is available at https://godoc.org/github.com/globalsign/mgo
//
// To get started, connect to MongoDB:
//
//     url := "mongodb://user:pass@host/db-name"
//     repo, err := mr.Connect(url)
//
// Include the "Base" type in your data types:
//
//     type User struct {
//             mr.Base `json:",inline" bson:",inline"`
//
//             Name         string `bson:"name" json:"name"`
//             Email        string `bson:"email" json:"email"`
//             PasswordHash []byte `bson:"password_hash" json:"-"`
//     }
//
// Doing so will add the following fields to your type:
//
//     ID: ObjectId of the document in the DB
//     CreatedAt: Timestamp of first insertion into DB
//     UpdatedAt: Timestamp of the last update operations
//
// The "MongoCollection" type implements the basic CRUD operations defined in
// the "Colleciton" interface. IDs can be passed as strings or bson.ObjectIDs,
// and objects are usually types that include the "Base" type (see above).
package mr
