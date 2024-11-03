package entity

type API struct {
	ID         int64
	Path       string
	Method     int8
	Host       string
	Scheme     string
	Field      string
	IsActive   int8 `gorm:"default:1"`
	ExtraData  string
	CreateTime int64 `gorm:"default:0"`
	UpdateTime int64 `gorm:"default:0"`
}

func (d *API) TableName() string {
	return "testing_api_tab"
}
