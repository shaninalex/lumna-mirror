package models

type Project struct {
	Id   uint
	Name string

	Boards    []*Board
	Calendars []*Calendar
	Documents *Folder
}

func (s *Project) GetId() uint  { return s.Id }
func (s *Project) SetId(u uint) { s.Id = u }
