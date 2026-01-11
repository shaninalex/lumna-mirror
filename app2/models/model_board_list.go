package models

type BoardList struct {
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	Order   uint   `json:"order"`
	BoardId uint   `json:"board_id"`
}

func (s *BoardList) GetId() uint  { return s.Id }
func (s *BoardList) SetId(u uint) { s.Id = u }
