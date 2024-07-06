// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameAssetgoods = "assetgoods"

// Assetgoods mapped from table <assetgoods>
type Assetgoods struct {
	ID         int64          `gorm:"column:id;not null;comment:资产表id" json:"id"`                                                            // 资产表id
	UID        int64          `gorm:"column:uid;primaryKey;comment:所得的人" json:"uid"`                                                         // 所得的人
	Goodsid    int64          `gorm:"column:goodsid;primaryKey;comment:商品id" json:"goodsid"`                                                 // 商品id
	Amount     int32          `gorm:"column:amount;comment:当前拥有数量" json:"amount"`                                                            // 当前拥有数量
	Spending   int32          `gorm:"column:spending;comment:已花费数量" json:"spending"`                                                         // 已花费数量
	Count      int64          `gorm:"column:count;comment:累计总数量" json:"count"`                                                               // 累计总数量
	Totalprice int64          `gorm:"column:totalprice;comment:总价(每次消费的购买)" json:"totalprice"`                                               // 总价(每次消费的购买)
	Code       int32          `gorm:"column:code;comment:最近一次的操作码(0:结算 1:充值 2:平台扣除 3:平台奖励 4:冻结 5:退税 6:提取 7:购买房卡 8:消耗房卡 9:置换房卡)" json:"code"` // 最近一次的操作码(0:结算 1:充值 2:平台扣除 3:平台奖励 4:冻结 5:退税 6:提取 7:购买房卡 8:消耗房卡 9:置换房卡)
	Time       int64          `gorm:"column:time;comment:最近一次消费记录" json:"time"`                                                              // 最近一次消费记录
	Remark     string         `gorm:"column:remark;comment:备注" json:"remark"`                                                                // 备注
	CreatedAt  time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	UpdateBy   int64          `gorm:"column:update_by" json:"update_by"`
	CreateBy   int64          `gorm:"column:create_by" json:"create_by"`
}

// TableName Assetgoods's table name
func (*Assetgoods) TableName() string {
	return TableNameAssetgoods
}