package service

// swagger:parameters signinHandler signupHandler
type User struct {
	// Existing username for /signin or new username for /signup
	//
	// required: true
	// in: body
	Username string `json:"username"`

	// Permissions group for new user in /signup, not required in /signin
	//
	// in: body
	Group string `json:"group"`

	// Bearer token containing password JWT
	//
	// required: true
	// in: header
	Password string `json:"password"`
}
