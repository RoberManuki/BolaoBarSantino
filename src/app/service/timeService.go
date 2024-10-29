package service

import (
	"bolao/src/app/model"
	repository "bolao/src/app/resource"
)

func GetTimes() ([]model.Time, error) {
	return repository.GetTimes()
}
