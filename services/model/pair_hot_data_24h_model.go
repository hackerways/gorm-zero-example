package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ PairHotData24HModel = (*customPairHotData24HModel)(nil)

type (
	// PairHotData24HModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPairHotData24HModel.
	PairHotData24HModel interface {
		pairHotData24HModel
		customPairHotData24HLogicModel
	}

	customPairHotData24HModel struct {
		*defaultPairHotData24HModel
	}

	customPairHotData24HLogicModel interface {
	}
)

// NewPairHotData24HModel returns a model for the database table.
func NewPairHotData24HModel(conn *gorm.DB, c cache.CacheConf) PairHotData24HModel {
	return &customPairHotData24HModel{
		defaultPairHotData24HModel: newPairHotData24HModel(conn, c),
	}
}

func (m *defaultPairHotData24HModel) customCacheKeys(data *PairHotData24H) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
