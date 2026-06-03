package svc

import (
	"context"
	"database/sql"
	"errors"

	db "github.com/cksidharthan/ghost-send/db/sqlc"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

var (
	ErrInvalidPassword = errors.New("invalid password")
	ErrSecretNotFound  = errors.New("secret not found or expired")
)

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

func (s *Service) GetSecret(c context.Context, request db.GetSecretByIDLockedParams) (*db.GetSecretByIDLockedRow, error) {
	s.Logger.Info("getting secret")
	secret, err := s.Store.AccessSecretAtomic(c, request)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrSecretNotFound
		}
		s.Logger.Error("error getting secret", zap.Error(err))
		return nil, err
	}

	if !secret.PasswordMatches {
		s.Logger.Error("password does not match")
		return nil, ErrInvalidPassword
	}

	return secret, nil
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
