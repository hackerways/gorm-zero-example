package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ PairHotData6HModel = (*customPairHotData6HModel)(nil)

type (
	// PairHotData6HModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPairHotData6HModel.
	PairHotData6HModel interface {
		pairHotData6HModel
		customPairHotData6HLogicModel
	}

	customPairHotData6HModel struct {
		*defaultPairHotData6HModel
	}

	customPairHotData6HLogicModel interface {
	}
)

// NewPairHotData6HModel returns a model for the database table.
func NewPairHotData6HModel(conn *gorm.DB, c cache.CacheConf) PairHotData6HModel {
	return &customPairHotData6HModel{
		defaultPairHotData6HModel: newPairHotData6HModel(conn, c),
	}
}

func (m *defaultPairHotData6HModel) customCacheKeys(data *PairHotData6H) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
