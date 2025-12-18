package models

type Folder struct {
	Id        uint
	Name      string
	Documents []*Document
	ParentId  *uint
	Children  []*Folder
}

func (s *Folder) GetId() uint  { return s.Id }
func (s *Folder) SetId(u uint) { s.Id = u }

type Document struct {
	Id       uint
	Name     string
	FolderId uint
}

func (s *Document) GetId() uint  { return s.Id }
func (s *Document) SetId(u uint) { s.Id = u }
