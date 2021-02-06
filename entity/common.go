package entity

import "time"

type YesOrNo int8
type Sex int8
type TradeState int8  //交易状态
type CommonState int8 //通用状态 ，正常/停用

const (
	Yes YesOrNo = 1
	No  YesOrNo = 2
)
const (
	Male   Sex = 1
	Female Sex = 2
)

type TraceTime struct {
	Created time.Time `json:"created" xorm:"created"`
	Updated time.Time `json:"updated" xorm:"updated"`
}
type PagerForm struct {
	Page  int `json:"page" form:"page"`
	Limit int `json:"limit" form:"limit"`
}

const (
	CommonStateOK  CommonState = 1  //正常
	CommonStateDel CommonState = -1 //停用
)

var CommonStates = map[CommonState]string{
	CommonStateOK:  "正常",
	CommonStateDel: "已停用",
}

func (u CommonState) String() string {
	if name, ok := CommonStates[u]; ok {
		return name
	} else {
		return "-"
	}
}

type IEntity interface {
	TableName() string
}
