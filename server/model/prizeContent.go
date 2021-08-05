package model
//奖品列表
type PrizeContent struct {
	PrizeId int64 `json:"prizeid" xorm:"not null pk unique"`
	PrizeName string  `json:"prizename" xorm:"not null"`
	PrizeAmount int64 `json:"prizeamount" xorm:"not null"`
}
