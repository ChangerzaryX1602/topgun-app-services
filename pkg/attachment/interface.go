package attachment

import "top-gun-app-services/pkg/models"

type AttachmentService interface {
	CreateFile(AttachFile) error
	GetFile(int) (AttachFile, error)
	GetDatas(models.Paginate) ([]AttachFile, error)
	GetData(int) (AttachFile, error)
}
type AttachmentRepository interface {
	CreateFile(AttachFile) error
	GetFile(int) (AttachFile, error)
	GetDatas(models.Paginate) ([]AttachFile, error)
	GetData(int) (AttachFile, error)
}
