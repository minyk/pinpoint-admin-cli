package commands

import (
	"errors"
	"github.com/asaskevich/govalidator"
	"gopkg.in/alecthomas/kingpin.v3-unstable"
	"strconv"
)

func CheckDurationDays(app *kingpin.Application, element *kingpin.ParseElement, context *kingpin.ParseContext) error {
	days, err := strconv.ParseInt(*element.Value, 10, 64)
	if err != nil {
		return errors.New("duration is wrong")
	} else if days < 30 {
		return errors.New("duration is longer than 30")
	}
	return nil
}

func CheckURL(app *kingpin.Application, element *kingpin.ParseElement, context *kingpin.ParseContext) error {
	if !govalidator.IsURL(*element.Value) {
		return errors.New("pinpoint url is not valid")
	}
	return nil
}