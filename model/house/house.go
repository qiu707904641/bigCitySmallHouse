package house

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const CollectionName = "house"

type House struct {
	Id       primitive.ObjectID `bson:"_id,omitempty"` // mongo id
	UId      string             `bson:"uid"`           // 怕不同房源source_id重复导致数据覆盖, source + - + source_id
	SourceId string             `bson:"source_id"`     // 来源的房子id
	Source   Source             `bson:"source"`        // 来源
	Type     Type               `bson:"type"`          // 房子类型， 公寓/住宅
	// 需要加个复式吗？
	Name        string      `bson:"name"`        // 小区名字/公寓名字/...
	Description string      `bson:"description"` // 描述
	ImgUrls     []string    `bson:"img_urls"`    // 图片
	VideoUrls   []string    `bson:"video_urls"`  // 视频
	Area        float64     `bson:"area"`        // 面积 单位/m²
	Price       Price       `bson:"price"`       // 价格
	Floor       int         `bson:"floor"`       // 楼层
	Location    Location    `bson:"location"`    // 地点
	RentType    RentType    `bson:"rentType"`    // 租住类型, 合租/整租
	BuildTime   time.Time   `bson:"build_time"`  // 建造日期
	Facilities  []string    `bson:"facility"`    // 设施   床、桌子、电梯、跑步机...
	Traffic     []Traffic   `bson:"traffic"`     // 交通
	Composition Composition `bson:"composition"` // n厅n房...组成
	UpdateAt    time.Time   `bson:"update_at"`   // 数据更新时间
}

const (
	TypeUnknown   Type = "unknown"   // 未知
	TypeApartment Type = "apartment" // 公寓
	TypeResidence Type = "residence" // 住宅
	TypeVilla     Type = "villa"     // 别墅
	TypeShop      Type = "shop"      // 商铺
	TypeParking   Type = "parking"   // 停车位
	TypeOffice    Type = "office"    // 办公楼
)

type Type string

func (receiver *Type) String() string {
	return string(*receiver)
}

type RentType string

const (
	RentTypeEntire = "entire" // 整租
	RentTypeShared = "shared" // 合租
)
