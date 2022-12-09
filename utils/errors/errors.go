package errors

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
	OpenFileError
	AppendFileError
	UnmarshalWrong

	CertFileNotExists
	KeyFileNotExists

	UserNotExists
	UserAlreadyExists
	UploadAvatarError
	BadPassword
	EmailFormatInvalid
	PictureFormatInvalid
	NotAcceptAttachment

	UserIsNotLogIn
	CredentialError
	HasNoPermission

	TokenIsInvalid
	ParseClaimError
	UnexpectedSigningMethod
	GenerateTokenError
	TokenHasExpired

	AccessTimeOut

	BindPostDataError
	HasDirtyWord
	PostTooFast

	ErrRecordNotFound

	HashMarshalError
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
	OpenFileError:     "file cannot opened",
	AppendFileError:   "file cannot appended",
	UnmarshalWrong:    "the file cannot unmarshaled",
	CertFileNotExists: "the certification is not exists",
	KeyFileNotExists:  "the key is not exists",

	UserNotExists:        "user is not exist",
	UserAlreadyExists:    "username or user mail is not unique",
	UploadAvatarError:    "upload avatar failed",
	BadPassword:          "user input password incorrect",
	EmailFormatInvalid:   "user's email format is invalid",
	PictureFormatInvalid: "the picture format is invalid",
	NotAcceptAttachment:  "this is a not accpetted attachement",

	UserIsNotLogIn:          "user is not log in",
	CredentialError:         "wrong email/username or password",
	HasNoPermission:         "you have no permission to access the resource",
	TokenIsInvalid:          "Token string is invalid",
	ParseClaimError:         "cloud not parse the claim from the token string",
	UnexpectedSigningMethod: "unexpected signing method",
	GenerateTokenError:      "cannot generate token",
	TokenHasExpired:         "Token has expired, please login again",
	AccessTimeOut:           "access time out",

	BindPostDataError: "the engine can not bind user post data",
	HasDirtyWord:      "has the dirty or sensitive word",
	PostTooFast:       "you are post too fast, please wait a moment and try again",

	ErrRecordNotFound: "the record is not found",

	HashMarshalError: "Hash marshal string error",
}
