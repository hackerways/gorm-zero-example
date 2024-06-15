package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ EvmDataModel = (*customEvmDataModel)(nil)

type (
	// EvmDataModel is an interface to be customized, add more methods here,
	// and implement the added methods in customEvmDataModel.
	EvmDataModel interface {
		evmDataModel
		customEvmDataLogicModel
	}

	customEvmDataModel struct {
		*defaultEvmDataModel
	}

	customEvmDataLogicModel interface {
	}
)

// NewEvmDataModel returns a model for the database table.
func NewEvmDataModel(conn *gorm.DB, c cache.CacheConf) EvmDataModel {
	return &customEvmDataModel{
		defaultEvmDataModel: newEvmDataModel(conn, c),
	}
}

func (m *defaultEvmDataModel) customCacheKeys(data *EvmData) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
