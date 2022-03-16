package model

type File struct {
	ID         uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	UploadUser uint   `json:"uploadUser"`
	URL        string `json:"URL"`
	//AllowedUser
}
