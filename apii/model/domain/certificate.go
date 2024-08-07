package domain

import "time"

type Certificate struct {
	Id         int    `gorm:"primaryKey"`
	UserName   string `gorm:"not null"`
	UserId     int
	CourseId   int
	CourseName string    `gorm:"not null"`
	Date       time.Time `gorm:"type:TIMESTAMP"`
	CertType   string
	ImageUri   string
	CertUri    string
}
