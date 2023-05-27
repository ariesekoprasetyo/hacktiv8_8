package db

import "time"

type WindWaterStatus struct {
	Id         uint      `gorm:"primaryKey" json:"id"`
	Wind       int       `json:"wind"`
	Water      int       `json:"water"`
	Updated_at time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
