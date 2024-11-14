package attachment

import "top-gun-app-services/pkg/models"

type AttachmentService interface {
	CreateFile(AttachFile) (AttachFile, error)
	GetFile(int) (AttachFile, error)
	GetDatas(models.Paginate) ([]AttachFile, error)
	GetData(int) (AttachFile, error)
}
type AttachmentRepository interface {
	CreateFile(AttachFile) (AttachFile, error)
	GetFile(int) (AttachFile, error)
	GetDatas(models.Paginate) ([]AttachFile, error)
	GetData(int) (AttachFile, error)
}
