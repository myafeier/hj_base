package entity

import "time"

var OrderProjectTableName string = "hj_order_project"

//订单项目记录
type OrderProject struct {
	Id               int64     `json:"id"`
	OrderId          int64     `json:"order_id" xorm:"default 0 index"`                        //商城订单对应着商城id,定向订单对应着定向订单logId
	Source           Source    `json:"source" xorm:"default 0 index"`                          //订单来源
	ChannelId        int64     `json:"channel_id" xorm:"default 0 index"`                      //渠道id
	DetectionId      int64     `json:"detection_id" xorm:"default 0 index"`                    //检验所id
	UserId           int64     `json:"user_id" xorm:"default 0 index"`                         //订单用户id
	ExamineeId       int64     `json:"examinee_id" xorm:"default 0 index"`                     //受检人id,现场采样订单改为1对多，弃用此字段
	GoodsId          int64     `json:"goods_id" xorm:"default 0 index"`                        //商品Id
	GoodsName        string    `json:"goods_name" xorm:"varchar(100) default ''"`              //商品名称
	GoodsNameAbbr    string    `json:"goods_name_abbr" xorm:"varchar(100) default ''"`         //商品名称
	GoodsNo          string    `json:"goods_no" xorm:"varchar(100) default '' index"`          //商品编码
	SamplerType      string    `json:"sampler_type" xorm:"default 0"`                          //采样器类型
	CollectorBarcode string    `json:"collector_barcode" xorm:"varchar(100) default '' index"` //采集器条码
	Projects         []string  `json:"projects" xorm:"json"`                                   //订单项目对应的检测项目编码集合
	Report           string    `json:"report" xorm:"varchar(200)"`                             //报告存储路径
	ScanStat         YesOrNo   `json:"scan_stat" xorm:"tinyint(2) default 0 "`                 //检验所扫码状态(发生在检索人员收到样本后的扫码入库操作) ,1 已读取
	BindDate         time.Time `json:"bind_date" xorm:""`                                      // 用户绑定时间
	ByUserId         int64     `json:"by_user_id" xorm:"default 0 index"`                      //添加人id
	BindDateF        string    `json:"bind_date_f" xorm:"-"`
	CreatedF         string    `json:"created_f" xorm:"-"`
	TraceTime        `xorm:"extends"`
}

func (e *OrderProject) TableName() string {
	return OrderProjectTableName
}
