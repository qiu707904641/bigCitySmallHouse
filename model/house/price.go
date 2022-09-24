package house

type Price struct {
	Rent                 float64 `bson:"rent"`                   // 租金
	ElectricityPerDegree float64 `bson:"electricity_per_degree"` // 电价/度
	WaterPerCube         float64 `bson:"water_per_cube"`         // 水价/立方
	ManagementPerMeter   float64 `bson:"management_per_meter"`   // 管理费/米
	ManagementTotal      float64 `bson:"management_total"`       // 管理费，一次性
}
