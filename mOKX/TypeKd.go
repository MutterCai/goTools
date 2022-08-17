package mOKX

import "time"

type CandleDataType [7]string

type TypeKd struct {
	InstID   string    `json:"InstID"`   // 持仓币种
	CcyName  string    `json:"CcyName"`  // 币种名称
	TimeUnix int64     `json:"TimeUnix"` // 毫秒时间戳
	Time     time.Time `json:"Time"`     // 时间
	O        string    `json:"O"`        // 开盘
	H        string    `json:"H"`        // 最高
	L        string    `json:"L"`        // 最低
	C        string    `json:"C"`        // 收盘价格
	CBas     string    `json:"CBas"`     // 实体中心价 (收盘+最高+最低) / 4
	Vol      string    `json:"Vol"`      // 交易货币的数量
	VolCcy   string    `json:"VolCcy"`   // 计价货币数量
	Type     string    `json:"Type"`     // 数据类型
	Dir      int       `json:"Dir"`      // 方向 (收盘-开盘) ，1：涨 & -1：跌 & 0：横盘
	HLPer    string    `json:"HLPer"`    // 振幅 (最高-最低)/最低 * 100%
	U_shade  string    `json:"U_shade"`  // 上影线
	D_shade  string    `json:"D_shade"`  // 下影线
	RosePer  string    `json:"RosePer"`  // 涨幅 当前收盘价 - 上一位收盘价 * 100%
	C_dir    int       `json:"C_dir"`    // 中心点方向 (当前中心点-前中心点) 1：涨 & -1：跌 & 0：横盘
	TickSz   string    `json:"tickSz"`   // 价格精度
}

// 基于 K线数据分析结果
type AnalySliceType struct {
	InstID        string    `json:"InstID"`    // InstID
	CcyName       string    `json:"CcyName"`   // 币种名称
	StartTime     time.Time `json:"StartTime"` // 开始时间
	StartTimeUnix int64     `json:"StartTimeUnix"`
	EndTime       time.Time `json:"EndTime"` // 结束时间
	EndTimeUnix   int64     `json:"EndTimeUnix"`
	DiffHour      int       `json:"DiffHour"`   // 总时长
	Len           int       `json:"Len"`        // 数据的总长度
	Volume        string    `json:"Volume"`     // 成交量总和
	VolumeAvg     string    `json:"VolumeAvg"`  // 平均 小时 成交量
	RosePer       string    `json:"RosePer"`    // 涨幅
	H             string    `json:"H"`          // 最高价
	L             string    `json:"L"`          // 最低价
	U_shadeAvg    string    `json:"U_shadeAvg"` // 上影线平均长度
	D_shadeAvg    string    `json:"D_shadeAvg"` // 下影线平均长度
	HLPerMax      string    `json:"HLPerMax"`   // 最高振幅
	HLPerAvg      string    `json:"HLPerAvg"`   // 平均振幅
	U_RIdx        int       `json:"U_RIdx"`     // 涨幅 = (最新价-开盘价)/开盘价 =
	VolIdx        int       `json:"VolIdx"`     // 成交量排名
}

type AnalySingleType struct {
	InstID        string    `json:"InstID"`    // InstID
	StartTime     time.Time `json:"StartTime"` // 开始时间
	StartTimeUnix int64     `json:"StartTimeUnix"`
	EndTime       time.Time `json:"EndTime"` // 结束时间
	EndTimeUnix   int64     `json:"EndTimeUnix"`
	DiffHour      int64     `json:"DiffHour"` // 总时长
}
