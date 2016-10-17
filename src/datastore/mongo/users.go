package mongo

import (
	"errors"

	"github.com/parkn-co/parkn-server/src/services/authentication"
	"github.com/parkn-co/parkn-server/src/types"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// UsersCollection is the struct we use to interact with the users collection
type UsersCollection struct {
	baseCollection
}

// NewUsersCollection creates a new UsersCollection
func NewUsersCollection(collection *mgo.Collection) *UsersCollection {
	return &UsersCollection{baseCollection{collection}}
}

// CreateUser creates a user object
func (users *UsersCollection) CreateUser(newUser *types.NewUser) (bson.ObjectId, error) {
	hash, err := authentication.PasswordHash(newUser.Password)
	if err != nil {
		return "", err
	}

	id := bson.NewObjectId()
	err = users.collection.Insert(types.User{
		ID:           id,
		Email:        newUser.Email,
		FirstName:    newUser.FirstName,
		LastName:     newUser.LastName,
		PasswordHash: hash,
	})

	return id, err
}

// UserExistsByEmail creates a user object
func (users *UsersCollection) UserExistsByEmail(email string) bool {
	count, _ := users.collection.Find(bson.M{"email": email}).Count()

	return count > 0
}

// GetUserByLogin gets a user by login request. It checks the validity of the password and email
func (users *UsersCollection) GetUserByLogin(loginRequest *types.LoginRequest, user *types.User) error {
	err := users.collection.Find(bson.M{"email": loginRequest.Email}).One(user)
	if err != nil {
		return err
	}

	if !authentication.Authenticate(user, loginRequest) {
		return errors.New("Not found")
	}

	return nil
}
