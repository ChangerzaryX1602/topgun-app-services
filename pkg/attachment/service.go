package attachment

import "top-gun-app-services/pkg/models"

type attachmentService struct {
	repo AttachmentRepository
}

func NewAttachmentService(repo AttachmentRepository) AttachmentService {
	return &attachmentService{repo}
}

func (s attachmentService) CreateFile(attach AttachFile) (AttachFile, error) {
	return s.repo.CreateFile(attach)
}
func (s attachmentService) GetFile(id int) (AttachFile, error) {
	return s.repo.GetFile(id)
}
func (s attachmentService) GetDatas(paginate models.Paginate) ([]AttachFile, error) {
	return s.repo.GetDatas(paginate)
}
func (s attachmentService) GetData(id int) (AttachFile, error) {
	return s.repo.GetFile(id)
}
