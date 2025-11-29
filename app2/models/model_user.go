package models

type User struct {
	Id     uint
	Email  string
	Active bool
}

func (s *User) GetId() uint       { return s.Id }
func (s *User) SetId(v uint)      { s.Id = v }
func (s *User) GetEmail() string  { return s.Email }
func (s *User) SetEmail(v string) { s.Email = v }
func (s *User) IsActive() bool    { return s.Active }
func (s *User) SetActive(v bool)  { s.Active = v }
