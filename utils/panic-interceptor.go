package utils

import "fmt"

func PanicInterceptor(outputError *error) {
	if err := recover(); err != nil {
		*outputError = fmt.Errorf("panic occurred: %v", err)
	}
}
