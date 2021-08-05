package model
//抽奖情况
type WheelPrize struct {
	WheelId    int64  `json:"wheelid" xorm:"not null pk autoincr"` //指定主键并自增
	PrizeName string  `json:"prizename" xorm:"not null"`
}

