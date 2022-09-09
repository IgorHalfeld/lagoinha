package service

import "github.com/igorhalfeld/lagoinha/internal/entity"

type Provider interface {
	Request(cep string) (*entity.Cep, error)
}
