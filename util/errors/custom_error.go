package errors

import (
	"fmt"
	"net/http"
)

type customError struct {
	message string
	code    int
}

func (ce *customError) Error() string {
	return ce.message
}

func (ce *customError) StatusCode() int {
	return ce.code
}

func newError(message string, code int) *customError {
	return &customError{
		message: message,
		code:    code,
	}
}

//Error Declaration
var (
	InvalidSigningAlgorithm           = invalidSigningAlgorithm{}
	InvalidAccessToken                = invalidAccessToken{}
	TransactionRollbackeNotExistError = transactionRollbackeNotExistError{}
)

// Error signing algorithm
type invalidSigningAlgorithm struct{}

func (invalidSigningAlgorithm) Error() string {
	return "Invalid signing algorithm"
}

func (invalidSigningAlgorithm) StatusCode() int {
	return http.StatusInternalServerError
}

// Error invalid access token
type invalidAccessToken struct{}

func (invalidAccessToken) Error() string {
	return "Invalid access token"
}

func (invalidAccessToken) StatusCode() int {
	return http.StatusBadRequest
}

//
type transactionRollbackeNotExistError struct{}

func (transactionRollbackeNotExistError) Error() string {
	return "Transaction rollback not exist"
}

func (transactionRollbackeNotExistError) StatusCode() int {
	return http.StatusInternalServerError
}

// Common
var (
	SomethingWrong = func(name string, value interface{}) *customError {
		return newError(fmt.Sprintf("Something wrong [%s=%v]", name, value), http.StatusInternalServerError)
	}
	TypeAssertionError    = newError("Type assertion error", http.StatusInternalServerError)
	InvalidParameterError = func(name string, value interface{}) *customError {
		return newError(fmt.Sprintf("Invalid parameter [%s=%v]", name, value), http.StatusBadRequest)
	}
	NoRecordsPassedToPerformOperationError = func(name string, value interface{}) *customError {
		return newError(fmt.Sprintf("No records passed to perform operation [%s=%v]", name, value), http.StatusBadRequest)
	}
	NoDBConnectionError   = newError("No DB connection error", http.StatusInternalServerError)
	MissingParameterError = func(name string) *customError {
		return newError(fmt.Sprintf("Missing parameter [%s]", name), http.StatusBadRequest)
	}
)

var (
	InvalidCAPEMKeyError = newError("Invalid CA PEM key", http.StatusInternalServerError)
	NotAcceptError       = newError("Not accept", http.StatusInternalServerError)
)

var (
	InvalidIDTypeError = newError("Invalid ID type", http.StatusBadRequest)
	NotLoggedInError   = newError("Please loggin to continue", http.StatusBadRequest)
)

// UUID
var (
	UUIDFromStringParsingError = func(name string, value interface{}) *customError {
		return newError(fmt.Sprintf("Cannot parse uuid from string [%s=%v]", name, value), http.StatusInternalServerError)
	}
)

// User
var (
	CreateUserFailedError  = newError("Create user failed", http.StatusInternalServerError)
	MissingNameUserError   = newError("Missing name of user", http.StatusBadRequest)
	MissingUserNameError   = newError("Missing user name", http.StatusBadRequest)
	MissingPasswordError   = newError("Missing pass word", http.StatusBadRequest)
	MissingRoleError       = newError("Missing role", http.StatusBadRequest)
	UsernameIsExistedError = newError("Username Is Existed", http.StatusBadRequest)
	UserNotFoundError      = newError("User No tFound Error", http.StatusNotFound)
	WrongPasswordError     = newError("Wrong Password Error", http.StatusBadRequest)
	WrongRoleError         = newError("Wrong Role Error", http.StatusBadRequest)
	LengthNameError        = newError("Length Name Error", http.StatusBadRequest)
	LengthUsernameError    = newError("Length Username Error", http.StatusBadRequest)
	LengthPasswordError    = newError("Length Password Error", http.StatusBadRequest)
	InvalidUserIDError     = newError("Invalid User ID Error", http.StatusBadRequest)
	UpdateUserFailedError  = newError("Update user failed", http.StatusInternalServerError)
	UserNotExistError      = newError("User Not Exist Error", http.StatusInternalServerError)
)

// Auth
var (
	AccountNotFoundError = newError("Account Not Found Error", http.StatusNotFound)
	AccessDeniedError    = newError("Access Denied Error", http.StatusForbidden)
)

//Hero
var (
	MissingNameHeroError   = newError("Missing Name Hero Error", http.StatusBadRequest)
	MissingHeroPowerError  = newError("Missing Hero Power Error", http.StatusBadRequest)
	LengthHeroPowerError   = newError("Length Hero Power Error", http.StatusBadRequest)
	LengthNameHeroError    = newError("Length Name Hero Error", http.StatusBadRequest)
	CreateHeroFailedError  = newError("Create hero failed", http.StatusInternalServerError)
	UpdateHeroFailedError  = newError("Update hero failed", http.StatusInternalServerError)
	HeronameIsExistedError = newError("Hero name Is Existed Error", http.StatusBadRequest)
	HeroNotExistError      = newError("Hero Not Exist Error", http.StatusBadRequest)
)

//Group
var (
	MissingNameGroupError   = newError("Missing Name Group Error", http.StatusBadRequest)
	MissingGroupPowerError  = newError("Missing Group Power Error", http.StatusBadRequest)
	LengthGroupPowerError   = newError("Length Group Power Error", http.StatusBadRequest)
	LengthNameGroupError    = newError("Length Name Group Error", http.StatusBadRequest)
	CreateGroupFailedError  = newError("Create Group failed", http.StatusInternalServerError)
	UpdateGroupFailedError  = newError("Update Group failed", http.StatusInternalServerError)
	GroupnameIsExistedError = newError("Group name Is Existed Error", http.StatusBadRequest)
	GroupNotExistError      = newError("Group Not Exist Error", http.StatusBadRequest)
)

//Image
var (
	MissingImagesError = newError("Missing Images Error", http.StatusBadRequest)
	LengthImagesError  = newError("Length Images Error", http.StatusBadRequest)
)
