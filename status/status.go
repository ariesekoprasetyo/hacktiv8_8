package status

import "Assigment_8/db"

type StatusReq struct {
	Wind  int `json:"wind"`
	Water int `json:"water"`
}

func StatusUpdate(status StatusReq) error {
	var statusValue db.WindWaterStatus
	statusValue = db.WindWaterStatus{
		Wind:  status.Wind,
		Water: status.Water,
	}
	if err := db.DB.Model(&db.WindWaterStatus{}).Where("id = 1").Updates(statusValue).Error; err != nil {
		return err
	}
	return nil
}
