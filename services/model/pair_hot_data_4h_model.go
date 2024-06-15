package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ PairHotData4HModel = (*customPairHotData4HModel)(nil)

type (
	// PairHotData4HModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPairHotData4HModel.
	PairHotData4HModel interface {
		pairHotData4HModel
		customPairHotData4HLogicModel
	}

	customPairHotData4HModel struct {
		*defaultPairHotData4HModel
	}

	customPairHotData4HLogicModel interface {
	}
)

// NewPairHotData4HModel returns a model for the database table.
func NewPairHotData4HModel(conn *gorm.DB, c cache.CacheConf) PairHotData4HModel {
	return &customPairHotData4HModel{
		defaultPairHotData4HModel: newPairHotData4HModel(conn, c),
	}
}

func (m *defaultPairHotData4HModel) customCacheKeys(data *PairHotData4H) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
