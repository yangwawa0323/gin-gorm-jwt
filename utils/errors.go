package utils

type ErrorUint16 uint16

const (
	// Service
	_ ErrorUint16 = iota
	NotImplemented
	ServicePortIsUsed
	DialServicePortFailed
	PanicConfigFile
	DatabaseCanNotAccess
	DataIsTooLong

	NoMorePages
	FileNotExist
	UnmarshalWrong

	CertFileNotExists
	KeyFileNotExists

	UserNotExists
	UserAlreadyExists
	EmailFormatInvalid
	PictureFormatInvalid
	NotAcceptAttachment

	UserIsNotLogIn
	CredentialError
	HasNoPermission

	TokenIsInvalid
	ParseClaimError
	GenerateTokenError
	TokenHasExpired

	AccessTimeOut

	BindPostDataError
	HasDirtyWord
	PostTooFast

	ErrRecordNotFound
)

var Errors = map[ErrorUint16]string{
	NotImplemented:        "not implemented yet",
	ServicePortIsUsed:     "server port is used",
	DialServicePortFailed: "dailed service port is failed. Is the service started?",
	DatabaseCanNotAccess:  "the database can not access",

	PanicConfigFile:   "you have to set the conf/app.yaml file proplely",
	DataIsTooLong:     "the data is too loooooong",
	NoMorePages:       "this is the last page, no more pages",
	FileNotExist:      "file is not exists",
	UnmarshalWrong:    "the file cannot unmarshaled",
	CertFileNotExists: "the certification is not exists",
	KeyFileNotExists:  "the key is not exists",

	UserNotExists:        "user is not exist",
	UserAlreadyExists:    "username or user mail is not unique",
	EmailFormatInvalid:   "user's email format is invalid",
	PictureFormatInvalid: "the picture format is invalid",
	NotAcceptAttachment:  "this is a not accpetted attachement",

	UserIsNotLogIn:     "user is not log in",
	CredentialError:    "wrong email/username or password",
	HasNoPermission:    "you have no permission to access the resource",
	TokenIsInvalid:     "Token string is invalid",
	ParseClaimError:    "cloud not parse the claim from the token string",
	GenerateTokenError: "cannot generate token",
	TokenHasExpired:    "Token has expired, please login again",
	AccessTimeOut:      "access time out",

	BindPostDataError: "the engine can not bind user post data",
	HasDirtyWord:      "has the dirty or sensitive word",
	PostTooFast:       "you are post too fast, please wait a moment and try again",

	ErrRecordNotFound: "the record is not found",
}
