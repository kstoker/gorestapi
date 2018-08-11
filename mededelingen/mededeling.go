package mededelingen

import "time"

// Mededeling : mapping to table mededeling
type Mededeling struct {
	ID       int64     `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	Text     string    `gorm:"column:text" json:"text"`
	DateFrom time.Time `gorm:"column:datefrom" json:"datefrom"`
	DateTo   time.Time `gorm:"column:dateto" json:"dateto"`
}
