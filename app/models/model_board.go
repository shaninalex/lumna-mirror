package models

type Board struct {
	Id        uint   `json:"id"`
	Name      string `json:"name"`
	ProjectId uint   `json:"project_id"`

	Tasks []*Task
	Lists []*BoardList

	// settings contains:
	// - order of lists on frontend
	Settings *string `json:"settings"`
}

func (s *Board) GetId() uint  { return s.Id }
func (s *Board) SetId(u uint) { s.Id = u }
