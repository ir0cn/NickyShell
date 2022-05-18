package service

func Throw(msg string) {
	panic(msg)
}

func CheckError(err error, msg string) {
	if err != nil {
		Throw(msg)
	}
}
