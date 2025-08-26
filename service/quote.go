package service

import (
	"context"
	"gin-redis-shell/cache"
	"gin-redis-shell/database"
	"gin-redis-shell/dto"
)

type QuoteService struct {
	quoteModel *database.QuoteModel
	cacheModel *cache.Cache
}

func NewQuoteService() *QuoteService {
	return &QuoteService{
		quoteModel: database.NewQuoteModel(),
		cacheModel: cache.NewCache(),
	}
}

func (s *QuoteService) GetQuote(ctx context.Context, req *dto.QuoteReq) (resp *dto.QuoteReq, err error) {
	if req.Tag == "" || req.Name == "" {
		//	todo 随机捞一条
	}
	data, err := s.cacheModel.GetQuote(req)
	if err != nil {
		return nil, err
	}
	if data.Data.Content == "" {
		//	 后续接入ai模型
	}
	return
}
