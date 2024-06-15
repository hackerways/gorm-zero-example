package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ PairHotData30MModel = (*customPairHotData30MModel)(nil)

type (
	// PairHotData30MModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPairHotData30MModel.
	PairHotData30MModel interface {
		pairHotData30MModel
		customPairHotData30MLogicModel
	}

	customPairHotData30MModel struct {
		*defaultPairHotData30MModel
	}

	customPairHotData30MLogicModel interface {
	}
)

// NewPairHotData30MModel returns a model for the database table.
func NewPairHotData30MModel(conn *gorm.DB, c cache.CacheConf) PairHotData30MModel {
	return &customPairHotData30MModel{
		defaultPairHotData30MModel: newPairHotData30MModel(conn, c),
	}
}

func (m *defaultPairHotData30MModel) customCacheKeys(data *PairHotData30M) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
