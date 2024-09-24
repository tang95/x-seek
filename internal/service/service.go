package service

import (
	"context"
	"github.com/patrickmn/go-cache"
	"github.com/tang95/x-seek/config"
	"github.com/tang95/x-seek/internal/model"
	"go.uber.org/zap"
)

type Service struct {
	config       *config.Server
	logger       *zap.Logger
	incidentRepo model.IncidentRepo
	userRepo     model.UserRepo
	teamRepo     model.TeamRepo
	transaction  Transaction
	cache        *cache.Cache
}

type Transaction interface {
	InTx(context.Context, func(ctx context.Context) error) error
}

func NewService(config *config.Server, logger *zap.Logger,
	incidentRepo model.IncidentRepo,
	transaction Transaction,
	userRepo model.UserRepo,
	teamRepo model.TeamRepo,
) (*Service, error) {
	return &Service{
		config:       config,
		logger:       logger,
		incidentRepo: incidentRepo,
		transaction:  transaction,
		userRepo:     userRepo,
		teamRepo:     teamRepo,
		cache:        cache.New(cache.NoExpiration, cache.NoExpiration),
	}, nil
}
