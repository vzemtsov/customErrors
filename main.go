package customErrors

import (
	"fmt"
	"time"
)

type error interface {
	Error() string
}

type CustomError struct {
	PackageName string
	Functions   string
	Message     string
	Err         error
}

func (e *CustomError) Error() string {
	dt := time.Now()

	errorString := fmt.Sprintf("%v. Package: %s; Function: %s; Message: %s; Error: %s", dt.Format(time.UnixDate), e.PackageName, e.Functions, e.Message, e.Err.Error())
	return errorString
}
