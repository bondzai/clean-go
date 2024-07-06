package services

import (
	"clean-go/internal/utils/errs"
	"log"
)

func handleDbError(err error) error {
	if err.Error() == "event already exists" {
		return errs.NewBadRequestError("event already exists")
	}

	if err.Error() == "event not found" {
		return errs.NewBadRequestError("event not found")
	}

	log.Println(err)
	return errs.NewUnExpectedError()
}
