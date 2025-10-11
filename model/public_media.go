package models

// PublicMedia represents publicly accessible media files (images, audio, video, documents)
type PublicMedia struct {
	BaseModel
	UUID          string  `gorm:"uniqueIndex;not null"` // UUID for file identification
	OriginalName  string  `gorm:"not null"`             // Original filename
	FileType      string  `gorm:"index;not null"`       // File category: img, audio, video, docs
	FileExtension string  `gorm:"not null"`             // File extension: jpg, png, mp3, mp4, pdf, etc.
	FileSize      int64   `gorm:"not null"`             // File size in bytes
	MimeType      string  `gorm:"not null"`             // MIME type
	FilePath      string  `gorm:"not null"`             // Full MinIO path (e.g., /public/img/uuid.jpg)
	UploadedBy    *string `gorm:"default:null"`         // User ID who uploaded
	Description   *string `gorm:"default:null"`         // Optional description
}

// TableName specifies the table name for PublicMedia
func (PublicMedia) TableName() string {
	return "public_media"
}
