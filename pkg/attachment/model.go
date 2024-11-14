package attachment

import "time"

type AttachFile struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	FileName  string    `json:"file_name"`
	FilePath  string    `json:"file_path"`
	FileType  string    `json:"file_type"`
	CreatedAt time.Time `json:"created_at"`
}
