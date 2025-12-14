package models

type Board struct {
	Id    uint
	Tasks []*Task
	Lists []*BoardList

	// settings contains:
	// - order of lists on frontend
	Settings string
}

func (s *Board) GetId() uint  { return s.Id }
func (s *Board) SetId(u uint) { s.Id = u }
