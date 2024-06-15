package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ PairHotData1HModel = (*customPairHotData1HModel)(nil)

type (
	// PairHotData1HModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPairHotData1HModel.
	PairHotData1HModel interface {
		pairHotData1HModel
		customPairHotData1HLogicModel
	}

	customPairHotData1HModel struct {
		*defaultPairHotData1HModel
	}

	customPairHotData1HLogicModel interface {
	}
)

// NewPairHotData1HModel returns a model for the database table.
func NewPairHotData1HModel(conn *gorm.DB, c cache.CacheConf) PairHotData1HModel {
	return &customPairHotData1HModel{
		defaultPairHotData1HModel: newPairHotData1HModel(conn, c),
	}
}

func (m *defaultPairHotData1HModel) customCacheKeys(data *PairHotData1H) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
