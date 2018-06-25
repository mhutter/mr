# MongoRepo

[![GoDoc](https://godoc.org/github.com/mhutter/mr?status.svg)](https://godoc.org/github.com/mhutter/mr)

Library to ease interaction with MongoDB.

This package is an opinionated abstraction for [mgo][].

## Getting Started

Connect to MongoDB (uses `$MONGODB_URI` from environment if set, or `localhost`):

```go
repo := mr.MustAutoconnect("db-name")
```

Include the "Base" type in your data types:

```go
type User struct {
        mr.Base `json:",inline" bson:",inline"`

        Name         string `bson:"name" json:"name"`
        Email        string `bson:"email" json:"email"`
        PasswordHash []byte `bson:"password_hash" json:"-"`
}
```

Doing so will add the following fields to your type:

* ID: ObjectId of the document in the DB
* CreatedAt: Timestamp of first insertion into DB
* UpdatedAt: Timestamp of the last update operations

The "MongoCollection" type implements the basic CRUD operations defined in the "Colleciton" interface. IDs can be passed as strings or bson.ObjectIDs, and objects are usually types that include the "Base" type (see above).

```go
var u User
err := repo.C("users").Find("5b2ec2bc5a39251ef18064f8", &u)
```

## HTTP Middleware

You can use `mr.Middleware(repo)` to create a middleware that injects the given Repository in each request:

```go
withDB := mr.Middleware(repo)
http.HandleFunc("/api/v1/users", withDB(myHandler))
```


> [Manuel Hutter](https://hutter.io) -
> GitHub [@mhutter](https://github.com/mhutter) -
> Twitter [@dratir](https://twitter.com/dratir)

[mgo]: https://github.com/globalsign/mgo
