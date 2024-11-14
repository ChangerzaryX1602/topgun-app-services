package attachment

import (
	"top-gun-app-services/pkg/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type attachmentRepository struct {
	db *gorm.DB
}

func NewAttachmentRepository(db *gorm.DB) AttachmentRepository {
	return &attachmentRepository{db}
}
func (r attachmentRepository) GetDatas(paginate models.Paginate) ([]AttachFile, error) {
	var attach []AttachFile
	err := r.db.Limit(paginate.Limit).Offset(paginate.Offset).Find(&attach).Error
	if err != nil {
		return []AttachFile{}, err
	}
	return attach, nil
}
func (r attachmentRepository) GetData(id int) (AttachFile, error) {
	var attach AttachFile
	err := r.db.Where("id = ?", id).First(&attach).Error
	if err != nil {
		return AttachFile{}, err
	}
	return attach, nil
}
func (r attachmentRepository) CreateFile(attach AttachFile) (AttachFile, error) {
	err := r.db.Preload(clause.Associations).Create(&attach)
	if err.Error != nil {
		return AttachFile{}, err.Error
	}
	return attach, nil
}
func (r attachmentRepository) GetFile(id int) (AttachFile, error) {
	var attach AttachFile
	err := r.db.Where("id = ?", id).First(&attach).Error
	if err != nil {
		return AttachFile{}, err
	}
	return attach, nil
}
