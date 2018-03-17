package datastore

type Tag struct {
	ID   string `json:"id"`
	Text string `gorm:"type:text" json:"text"`
}
