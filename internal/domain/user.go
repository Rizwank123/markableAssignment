package domain

import (
	"context"

	"github.com/gofrs/uuid/v5"
)

type (
	// UserRole represents the role a user in the system.
	UserRole string
)
type (
	// User defines the module for User
	User struct {
		Base
		UserName string  `db:"user_name" json:"user_name,omitempty" example:"+919984778491"`
		Password *string `db:"password" json:"-"`
		Role     string  `db:"role" json:"role,omitempty"  example:"RECEPTIONIST"`
		FullName string  `db:"full_name" json:"full_name,omitempty" example:"John Doe"`
		BaseAudit
	} // @name User
)

type (
	// CreateUserInput define the module for CreateUser
	RegisterUserInput struct {
		FullName string   `json:"full_name" example:"John Doe"`
		UserName string   `json:"user_name" example:"+919984778491"`
		Role     UserRole `json:"role" example:"RECEPTIONIST"`
		Password string   `json:"password" example:"password123"`
	} // @name CreateUserInput
	// UpdateUserInput define the module for the UpdateUserInput
	UpdateUserInput struct {
		RegisterUserInput
	} // @name UpdateUserInput
	// LoginInput  define the module for the LoginInput
	LoginInput struct {
		UserName string `json:"user_name" example:"+919984778491 or example"`
		Password string `json:"password"`
	} // @name LoginInput
	// LoginOutput define the module for the LoginOutput
	LoginOutput struct {
		Token     string `json:"token"`
		ExpiresIn int64  `json:"expires_in"`
	} // @name LoginOutput
)

type (
	// UserRepository defines the methods that any use repository should implements
	UserRepository interface {
		// FindByID return the user by id
		FindByID(ctx context.Context, id uuid.UUID) (result User, err error)
		// FindByUserName return the user by username
		FindByUserName(ctx context.Context, username string) (result User, err error)
		// CreateUser creates a new user
		CreateUser(ctx context.Context, entity *User) (err error)
		// UpdateUser updates the user
		UpdateUser(ctx context.Context, entity *User) (err error)
		// DeleteUser deletes the user
		DeleteUser(ctx context.Context, id uuid.UUID) (err error)
	} // @name UserRepository

	// UserService defines the methods that any use service should implements
	UserService interface {
		// CreateUser creates a new user
		Register(in RegisterUserInput) (result User, err error)
		// FindByID  return the user by id
		FindByID(id uuid.UUID) (result User, err error)
		// FindByUserName return the user by username
		FindByUserName(username string) (result User, err error)
		// Login return the user by username and password
		Login(in LoginInput) (result LoginOutput, err error)
		// UpdateUser updates the user
		UpdateUser(id uuid.UUID, in UpdateUserInput) (result User, err error)
		// DeleteUser deletes the user
		DeleteUser(id uuid.UUID) (err error)
	} // @name UserService

)

const (
	UserRoleDoctor       UserRole = "DOCTOR"
	UserRoleRECEPTIONIST UserRole = "RECEPTIONIST"
)
