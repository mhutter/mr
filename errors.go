package mr

import "fmt"

// ErrInvalidURL is returned when the given connection URL could not be parsed
type ErrInvalidURL string

func (e ErrInvalidURL) Error() string {
	return "Invalid MongoDB URL: " + string(e)
}

// ErrNoObjectID is returned when you try to use a method that expects an object
// id with an ID that is not a valid ObjectID. See `bson.IsObjectIdHex`.
type ErrNoObjectID string

func (e ErrNoObjectID) Error() string {
	return fmt.Sprintf("'%s' is no valid ObjectID", string(e))
}
