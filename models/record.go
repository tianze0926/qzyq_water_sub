package models

// Submit record
type Record struct {
	ID         uint   `json:"id"`
	Time       int    `json:"time" gorm:"autoCreateTime:milli"`
	IP         string `json:"ip" binding:"required"`
	FileNumber string `json:"file_number" binding:"required"`
}
