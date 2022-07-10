package outputdata

type User struct {
	ID      string
	Email   string
	Profile *Profile
}

// SignUp ...
type SignUp struct {
	Email   string
	Message string
}

// SignIn ...
type SignIn struct {
	User User
	JWT  string
}

// ForgotPassword ...
type ForgotPassword struct {
	Email   string
	Message string
}

// UserSimplified ...
type UserSimplified struct {
	ID   string
	Name string
}

type PublicUser struct {
	ID   string
	Name string
}
