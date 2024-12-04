// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameMultisigHistory = "multisig_history"

// MultisigHistory mapped from table <multisig_history>
type MultisigHistory struct {
	TxHash         string    `gorm:"column:tx_hash;primaryKey" json:"tx_hash"`
	Status         int32     `gorm:"column:status" json:"status"`
	CreateTime     time.Time `gorm:"column:create_time;not null" json:"create_time"`
	MultisigAddr   string    `gorm:"column:multisig_addr;not null" json:"multisig_addr"`
	Memo           string    `gorm:"column:memo" json:"memo"`
	Fees           string    `gorm:"column:fees" json:"fees"`
	Height         string    `gorm:"column:height" json:"height"`
	GasUsed        string    `gorm:"column:gas_used" json:"gas_used"`
	GasWanted      string    `gorm:"column:gas_wanted" json:"gas_wanted"`
	TotalAmount    string    `gorm:"column:total_amount" json:"total_amount"`
	SequenceNumber string    `gorm:"column:sequence_number" json:"sequence_number"`
	Tx             string    `gorm:"column:tx" json:"tx"`
}

// TableName MultisigHistory's table name
func (*MultisigHistory) TableName() string {
	return TableNameMultisigHistory
}
