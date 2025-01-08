package svc

import (
	"context"
	"errors"
	"strings"

	db "github.com/cksidharthan/share-secret/db/sqlc"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

var ErrInvalidPassword = errors.New("invalid password")

type Service struct {
	Logger *zap.SugaredLogger
	Store  *db.Store
}

func New(logger *zap.SugaredLogger, store *db.Store) Service {
	return Service{
		Logger: logger,
		Store:  store,
	}
}

func (s *Service) CreateSecret(c context.Context, request db.CreateSecretParams) (*uuid.UUID, error) {
	s.Logger.Info("creating secret")
	secret, err := s.Store.CreateSecret(c, request)
	if err != nil {
		return nil, err
	}

	return &secret.ID, nil
}

func (s *Service) GetSecret(c context.Context, request db.GetSecretByIDParams) (*db.GetSecretByIDRow, error) {
	s.Logger.Info("getting secret")
	secret, err := s.Store.GetSecretByID(c, request)
	if err != nil {
		if strings.Contains(err.Error(), "converting NULL to string is unsupported") {
			s.Logger.Error("error getting secret", zap.Error(err))
			return nil, ErrInvalidPassword
		}

		s.Logger.Error("error getting secret", zap.Error(err))
		return nil, err
	}

	if !secret.PasswordMatches {
		s.Logger.Error("password does not match")
		return nil, errors.New("password does not match")
	}

	_, decrementErr := s.Store.DecrementTries(c, request.SecretID)
	if decrementErr != nil {
		s.Logger.Error("error decrementing tries", zap.Error(decrementErr))
	}

	if secret.RemainingTries <= 1 {
		deleteErr := s.Store.DeleteSecret(c, request.SecretID)
		if deleteErr != nil {
			s.Logger.Error("error deleting secret", zap.Error(deleteErr))
		}
	}

	return &secret, nil
}

// CheckSecretExists checks if a secret exists in the database
func (s *Service) CheckSecretExists(c context.Context, request uuid.UUID) (bool, error) {
	s.Logger.Info("checking if secret exists")
	secret, err := s.Store.CheckSecretStatus(c, request)
	if err != nil {
		s.Logger.Error("error getting secret", zap.Error(err))
		return false, err
	}

	return secret, nil
}
