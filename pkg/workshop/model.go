package workshop

import "time"

type EnergyConsumption struct {
	Power float64 `json:"power"`
}

// Voltage stores voltage data for different phases
type Voltage struct {
	L1 float64 `json:"l1_gnd"`
	L2 float64 `json:"l2_gnd"`
	L3 float64 `json:"l3_gnd"`
}

// MachineData represents the structure of the machine data
type MachineData struct {
	ID                 uint    `gorm:"primaryKey"`
	Power              float64 `json:"-"` // Store directly from EnergyConsumption.Power
	VoltageL1          float64 `json:"-"` // Store directly from Voltage.L1
	VoltageL2          float64 `json:"-"` // Store directly from Voltage.L2
	VoltageL3          float64 `json:"-"` // Store directly from Voltage.L3
	Pressure           float64 `json:"pressure"`
	Force              float64 `json:"force"`
	CycleCount         int     `json:"cycle_count"`
	PositionOfThePunch float64 `json:"position_of_the_punch"`
}

type RawData struct {
	ID                int               `json:"id" gorm:"primaryKey;autoIncrement"`
	EnergyConsumption EnergyConsumption `json:"energy_consumption" gorm:"type:jsonb"`
	Voltage           Voltage           `json:"voltage"`
	Pressure          float64           `json:"pressure"`
	Force             float64           `json:"force"`
	CycleCount        int               `json:"cycle_count"`
	PositionOfPunch   float64           `json:"position_of_the_punch"`
	CreatedAt         time.Time         `json:"created_at"`
}
