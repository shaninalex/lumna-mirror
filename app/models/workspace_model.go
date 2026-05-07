package models

type WorkspaceModel struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}
