package model

type MyError struct {
	Err string `json:"error"`
}

func (me MyError) Error() string {
	return me.Err
}

var (
	ErrorInvalidEmailOrPassword = MyError{
		Err: "Invalid email or password!",
	}
	ErrorInvalidToken = MyError{
		Err: "Invalid token!",
	}

	ErrorNotAuthorized = MyError{
		Err: "Not Authorized!",
	}

	ErrorNotFound = MyError{
		Err: "Not Found!",
	}

	ErrorForbiddenAccess = MyError{
		Err: "Forbidden Access!",
	}

	ErrorPhotoNotFound = MyError{
		Err: "Photo not found!",
	}
)

// var (
// 	ErrorInvalidEmailOrPassword = errors.New("Invalid email or password!")
// 	ErrorInvalidToken           = errors.New("Invalid token!")
// 	ErrorNotAuthorized          = errors.New("Not authorized!")
// 	ErrorNotFound               = errors.New("Not ound!")
// 	ErrorNotAuthorized          = errors.New()
// 	ErrorNotAuthorized          = errors.New()
// 	ErrorNotAuthorized          = errors.New()
// 	ErrorNotAuthorized          = errors.New()
// 	ErrorNotAuthorized          = errors.New()
// 	ErrorNotAuthorized          = errors.New()
// 	ErrorNotAuthorized          = errors.New()
// )
