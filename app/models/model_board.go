package models

type Board struct {
	Id    uint
	Tasks []*Task
	Lists []*List

	// settings contains:
	// - order of lists on frontend
	Settings string
}

func (s *Board) GetId() uint  { return s.Id }
func (s *Board) SetId(u uint) { s.Id = u }

type List struct {
	Id   uint
	Name string
	// order of list in the board
	Order uint
}

func (s *List) GetId() uint  { return s.Id }
func (s *List) SetId(u uint) { s.Id = u }
