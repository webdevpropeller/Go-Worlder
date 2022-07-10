package inputdata

// User ...
type User struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,password"`
	Category int    `validate:"required,min=1,max=2"`
}

// SignUp ...
type SignUp struct {
	User *User
}

// SignIn ...
type SignIn struct {
	Email    string
	Password string
}

// ForgotPassword ...
type ForgotPassword struct {
	Email string
}

// ResetPassword ...
type ResetPassword struct {
	Token    string
	Password string
}
