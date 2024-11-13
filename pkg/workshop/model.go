package workshop

import "time"

type EnergyConsumption struct {
	Power float64 `json:"Power"`
}

// Voltage stores voltage data for different phases
type Voltage struct {
	L1 float64 `json:"L1-GND"`
	L2 float64 `json:"L2-GND"`
	L3 float64 `json:"L3-GND"`
}

// MachineData represents the structure of the machine data
type MachineData struct {
	ID                 uint      `gorm:"primaryKey"`
	Power              float64   `json:"-"` // Store directly from EnergyConsumption.Power
	VoltageL1          float64   `json:"-"` // Store directly from Voltage.L1
	VoltageL2          float64   `json:"-"` // Store directly from Voltage.L2
	VoltageL3          float64   `json:"-"` // Store directly from Voltage.L3
	Pressure           float64   `json:"Pressure"`
	Force              float64   `json:"Force"`
	CycleCount         int       `json:"Cycle Count"`
	PositionOfThePunch float64   `json:"Position of the Punch"`
	CreatedAt          time.Time `json:"-" gorm:"autoCreateTime"`
}

type RawData struct {
	EnergyConsumption EnergyConsumption `json:"Energy Consumption" gorm:"type:jsonb"`
	Voltage           Voltage           `json:"Voltage"`
	Pressure          float64           `json:"Pressure"`
	Force             float64           `json:"Force"`
	CycleCount        int               `json:"Cycle Count"`
	PositionOfPunch   float64           `json:"Position of the Punch"`
}
