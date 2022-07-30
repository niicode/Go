package greeting

import (
	"fmt"
	"errors"
)

func Hello(name string) (string , error) {

	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf("Hello, %s!", name)
	return message, nil
}

