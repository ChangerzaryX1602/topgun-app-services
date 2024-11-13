package workshop

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"
	"top-gun-app-services/pkg/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type workshopRepository struct {
	db *gorm.DB
}

func NewWorkshopRepository(db *gorm.DB) WorkshopRepository {
	return &workshopRepository{db}
}
func (v EnergyConsumption) Value() (driver.Value, error) {
	return json.Marshal(v)
}
func (v *EnergyConsumption) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("Invalid type")
	}
	return json.Unmarshal(bytes, v)
}
func (v Voltage) Value() (driver.Value, error) {
	return json.Marshal(v)
}

// Implement the sql.Scanner interface to retrieve Voltage from JSON in the database
func (v *Voltage) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("Invalid type for Voltage")
	}
	return json.Unmarshal(bytes, v)
}
func (r workshopRepository) ProcessMessage(message []byte) {
	var machine RawData
	// Parse JSON message
	err := json.Unmarshal(message, &machine)
	if err != nil {
		log.Printf("Error parsing JSON: %v", err)
		return
	}

	// Save to database
	err = r.db.Create(&machine).Error
	if err != nil {
		fmt.Println("cannot write data into db", machine)
	} else {
		fmt.Println("write to db data:", machine)
	}
}
func (r workshopRepository) CreateMachine(data RawData) (RawData, error) {
	data.CreatedAt = time.Now()
	err := r.db.Preload(clause.Associations).Create(&data).Error
	if err != nil {
		return RawData{}, err
	}
	return data, nil
}
func (r workshopRepository) GetMachines(paginate models.Paginate) ([]RawData, error) {
	var machines []RawData
	err := r.db.Preload(clause.Associations).Limit(paginate.Limit).Offset(paginate.Offset).Find(&machines).Error
	if err != nil {
		return nil, err
	}
	return machines, nil
}
func (r workshopRepository) GetMachine(id string) (RawData, error) {
	var machine RawData
	err := r.db.Model(RawData{}).Preload(clause.Associations).Where("id = ?", id).First(&machine).Error
	if err != nil {
		return RawData{}, err
	}
	return machine, nil
}
func (r workshopRepository) UpdateMachine(id string, data RawData) (RawData, error) {
	err := r.db.Model(&RawData{}).Preload(clause.Associations).Where("id = ?", id).Updates(data).Error
	if err != nil {
		return RawData{}, err
	}
	return data, err
}
func (r workshopRepository) DeleteMachine(id string) error {
	err := r.db.Where("id = ?", id).Delete(&RawData{}).Error
	if err != nil {
		return err
	}
	return nil
}
