package service

import (
	"errors"
	repository "test_backend_2/Repository"

	"github.com/opentracing/opentracing-go/log"
)

type Service struct {
	repository repository.AreaRepository
}

func NewService(repository *repository.AreaRepository) *Service {
	return &Service{repository: *repository}
}

func (au *Service) Insert() error {
	err := au.repository.InsertArea(10, 10, "persegi")
	if err != nil {
		// log.Error().Msg(err.Error())
		// err = errors.New(en.ERROR_DATABASE)

		log.Error(err)
		err = errors.New("en.ERROR_DATABASE")
	}
	return err
}
