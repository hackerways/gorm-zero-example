package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ PairHotData5MModel = (*customPairHotData5MModel)(nil)

type (
	// PairHotData5MModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPairHotData5MModel.
	PairHotData5MModel interface {
		pairHotData5MModel
		customPairHotData5MLogicModel
	}

	customPairHotData5MModel struct {
		*defaultPairHotData5MModel
	}

	customPairHotData5MLogicModel interface {
	}
)

// NewPairHotData5MModel returns a model for the database table.
func NewPairHotData5MModel(conn *gorm.DB, c cache.CacheConf) PairHotData5MModel {
	return &customPairHotData5MModel{
		defaultPairHotData5MModel: newPairHotData5MModel(conn, c),
	}
}

func (m *defaultPairHotData5MModel) customCacheKeys(data *PairHotData5M) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
