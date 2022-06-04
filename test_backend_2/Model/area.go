package Model

type (
	Area struct {
		ID        int64   `gorm:"column:id;primaryKeyautoIncrement:true;"`
		AreaValue float64 `gorm:"column:area_value"`
		AreaType  string  `gorm:"column:type"`
	}
)
