package service

import (
	"blogService/global"
	"blogService/internal/dao"
	"context"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{
		ctx: ctx,
		dao: dao.New(global.DBEngine),
	}
	return svc
}
