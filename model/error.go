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
)
