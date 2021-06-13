package util

import (
	"log"
	"strings"

	"github.com/fgunawan1995/bcg/model"
	"github.com/pkg/errors"
)

func FormatGQLError(err error) error {
	if err == nil {
		return nil
	}
	errCause := errors.Cause(err)
	if strings.Contains(errCause.Error(), model.OutputErrorIdentifier) {
		return errors.New(strings.ReplaceAll(errCause.Error(), model.OutputErrorIdentifier, ""))
	}
	log.Printf("error = %+v", err)
	return errors.New(model.DefaultErrorMessage)
}
