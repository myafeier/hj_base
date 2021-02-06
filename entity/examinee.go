package entity

import "time"

var ExamineeTableName string = "hj_examinee"

//受检人信息

type Examinee struct {
	Id        int64       `json:"id"`
	Source    Source      `json:"source" xorm:"default 0 index"`                //来源
	UserId    int64       `json:"user_id" xorm:"default  0 index"`              //商城用户id
	State     CommonState `json:"state" xorm:"default 0 index"`                 //状态
	Default   YesOrNo     `json:"default" xorm:"tinyint(2) default 0"`          //是否默认
	Sex       Sex         `json:"sex" xorm:"varchar(2) default ''"`             //性别
	Age       string      `json:"age" xorm:"varchar(20) default ''"`            //年龄
	Birthday  time.Time   `json:"birthday" xorm:""`                             //生日
	Name      string      `json:"name" xorm:"varchar(100) default '' index"`    //受检人姓名
	IdCard    string      `json:"id_card" xorm:"varchar(100) default '' index"` //身份证
	Phone     string      `json:"phone" xorm:"varchar(100) default '' index"`   //电话
	Nation    string      `json:"nation" xorm:"varchar(20) default ''"`         //民族
	Address   string      `json:"address" xorm:"varchar(180) default ''"`       //住址
	Picture   string      `json:"picture" xorm:"text"`                          //身份证照片
	TraceTime `xorm:"extends"`
}

func (e *Examinee) TableName() string {
	return ExamineeTableName
}
