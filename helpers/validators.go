package helpers

import (
	"fmt"
	"strconv"

	"github.com/AlecAivazis/survey/v2"
)

func InRange(seasons int64) survey.Validator {
	return func(val interface{}) error {
		value, ok := val.(string)
		if !ok {
			return fmt.Errorf("invalid value")
		}

		number, err := strconv.ParseInt(value, 10, 32)
		if err != nil {
			return fmt.Errorf("not a number")
		}

		if number <= seasons && number > 0 {
			return nil
		}
		return fmt.Errorf("value not in range")
	}
}

func IsInt(val interface{}) error {
	nval, ok := val.(string)
	if !ok {
		return fmt.Errorf("invalid value")
	}

	interger, err := strconv.ParseInt(nval, 10, 32)
	if err != nil {
		return err
	}

	if interger <= 0 {
		return fmt.Errorf("ID cant be negative")
	}

	return nil
}
