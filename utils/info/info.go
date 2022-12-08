package info

type InfoUint16 uint16

const (
	_ InfoUint16 = iota
	Login
	Logout
	Register
	ChangePassword
)

var Infos = map[InfoUint16]string{
	Login:          "user log in",
	Logout:         "user log out",
	Register:       "new user register",
	ChangePassword: "user change his/her password",
}
