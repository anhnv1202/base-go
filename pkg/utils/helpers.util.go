package utils

import "github.com/anhnv1202/base-go/global"

func CheckErrorPanic(err error, errString string) {
	if err != nil {
		global.Logger.Err(errString, err)
		panic(errString)
	}
}