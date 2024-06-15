// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"fmt"
	"time"

	"github.com/SpectatorNan/gorm-zero/gormc"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var (
	cacheDataFlowPairHotData5MIdPrefix                 = "cache:dataFlow:pairHotData5M:id:"
	cacheDataFlowPairHotData5MChainIdPairAddressPrefix = "cache:dataFlow:pairHotData5M:chainId:pairAddress:"
)

type (
	pairHotData5MModel interface {
		Insert(ctx context.Context, tx *gorm.DB, data *PairHotData5M) error

		FindOne(ctx context.Context, id int64) (*PairHotData5M, error)
		FindOneByChainIdPairAddress(ctx context.Context, chainId int64, pairAddress string) (*PairHotData5M, error)
		Update(ctx context.Context, tx *gorm.DB, data *PairHotData5M) error

		Delete(ctx context.Context, tx *gorm.DB, id int64) error
		Transaction(ctx context.Context, fn func(db *gorm.DB) error) error
	}

	defaultPairHotData5MModel struct {
		gormc.CachedConn
		table string
	}

	PairHotData5M struct {
		Id          int64     `gorm:"column:id"`
		ChainId     int64     `gorm:"column:chain_id"`     // 加密货币id
		PairAddress string    `gorm:"column:pair_address"` // 交易对地址
		UpDown      float64   `gorm:"column:up_down"`      // 涨跌幅
		Volume      float64   `gorm:"column:volume"`       // 交易量
		SellCount   int64     `gorm:"column:sell_count"`   // 卖出数
		BuyCount    int64     `gorm:"column:buy_count"`    // 买入数
		BuyNetUsd   float64   `gorm:"column:buy_net_usd"`  // 净买入
		CreateAt    time.Time `gorm:"column:create_at"`
		UpdateAt    time.Time `gorm:"column:update_at"`
	}
)

func (PairHotData5M) TableName() string {
	return "`pair_hot_data_5m`"
}

func newPairHotData5MModel(conn *gorm.DB, c cache.CacheConf) *defaultPairHotData5MModel {
	return &defaultPairHotData5MModel{
		CachedConn: gormc.NewConn(conn, c),
		table:      "`pair_hot_data_5m`",
	}
}

func (m *defaultPairHotData5MModel) Insert(ctx context.Context, tx *gorm.DB, data *PairHotData5M) error {

	err := m.ExecCtx(ctx, func(conn *gorm.DB) error {
		db := conn
		if tx != nil {
			db = tx
		}
		return db.Save(&data).Error
	}, m.getCacheKeys(data)...)
	return err
}

func (m *defaultPairHotData5MModel) FindOne(ctx context.Context, id int64) (*PairHotData5M, error) {
	dataFlowPairHotData5MIdKey := fmt.Sprintf("%s%v", cacheDataFlowPairHotData5MIdPrefix, id)
	var resp PairHotData5M
	err := m.QueryCtx(ctx, &resp, dataFlowPairHotData5MIdKey, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&PairHotData5M{}).Where("`id` = ?", id).First(&resp).Error
	})
	switch err {
	case nil:
		return &resp, nil
	case gormc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultPairHotData5MModel) FindOneByChainIdPairAddress(ctx context.Context, chainId int64, pairAddress string) (*PairHotData5M, error) {
	dataFlowPairHotData5MChainIdPairAddressKey := fmt.Sprintf("%s%v:%v", cacheDataFlowPairHotData5MChainIdPairAddressPrefix, chainId, pairAddress)
	var resp PairHotData5M
	err := m.QueryRowIndexCtx(ctx, &resp, dataFlowPairHotData5MChainIdPairAddressKey, m.formatPrimary, func(conn *gorm.DB, v interface{}) (interface{}, error) {
		if err := conn.Model(&PairHotData5M{}).Where("`chain_id` = ? and `pair_address` = ?", chainId, pairAddress).Take(&resp).Error; err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case gormc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultPairHotData5MModel) Update(ctx context.Context, tx *gorm.DB, data *PairHotData5M) error {
	old, err := m.FindOne(ctx, data.Id)
	if err != nil && err != ErrNotFound {
		return err
	}
	err = m.ExecCtx(ctx, func(conn *gorm.DB) error {
		db := conn
		if tx != nil {
			db = tx
		}
		return db.Save(data).Error
	}, m.getCacheKeys(old)...)
	return err
}

func (m *defaultPairHotData5MModel) getCacheKeys(data *PairHotData5M) []string {
	if data == nil {
		return []string{}
	}
	dataFlowPairHotData5MChainIdPairAddressKey := fmt.Sprintf("%s%v:%v", cacheDataFlowPairHotData5MChainIdPairAddressPrefix, data.ChainId, data.PairAddress)
	dataFlowPairHotData5MIdKey := fmt.Sprintf("%s%v", cacheDataFlowPairHotData5MIdPrefix, data.Id)
	cacheKeys := []string{
		dataFlowPairHotData5MChainIdPairAddressKey, dataFlowPairHotData5MIdKey,
	}
	cacheKeys = append(cacheKeys, m.customCacheKeys(data)...)
	return cacheKeys
}

func (m *defaultPairHotData5MModel) Delete(ctx context.Context, tx *gorm.DB, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		if err == ErrNotFound {
			return nil
		}
		return err
	}
	err = m.ExecCtx(ctx, func(conn *gorm.DB) error {
		db := conn
		if tx != nil {
			db = tx
		}
		return db.Delete(&PairHotData5M{}, id).Error
	}, m.getCacheKeys(data)...)
	return err
}

func (m *defaultPairHotData5MModel) Transaction(ctx context.Context, fn func(db *gorm.DB) error) error {
	return m.TransactCtx(ctx, fn)
}

func (m *defaultPairHotData5MModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheDataFlowPairHotData5MIdPrefix, primary)
}

func (m *defaultPairHotData5MModel) queryPrimary(conn *gorm.DB, v, primary interface{}) error {
	return conn.Model(&PairHotData5M{}).Where("`id` = ?", primary).Take(v).Error
}