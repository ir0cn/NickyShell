package exception

import "fmt"

/****************************************************
 * 异常处理
 ****************************************************/

// Throw 调用panic抛异常
func Throw(msg interface{}) {
	panic(msg)
}

// ThrowIfError 调用panic抛异常
func ThrowIfError(err error) {
	if err != nil {
		Throw(err)
	}
}

// ThrowIfErrorWithMessage 调用panic抛异常
func ThrowIfErrorWithMessage(err error, msg string) {
	if err != nil {
		Throw(fmt.Sprintf("%v: %v", msg, err))
	}
}
