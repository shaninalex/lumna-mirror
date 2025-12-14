package models

type BoardList struct {
	Id      uint
	Name    string
	Order   uint
	BoardId uint
}

func (s *BoardList) GetId() uint  { return s.Id }
func (s *BoardList) SetId(u uint) { s.Id = u }
