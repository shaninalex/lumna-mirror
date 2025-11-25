package models

import (
	"time"
)

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

type Task struct {
	Id        uint
	Name      string
	ListId    uint
	Done      bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (s *Task) GetId() uint             { return s.Id }
func (s *Task) SetId(u uint)            { s.Id = u }
func (s *Task) GetCreatedAt() time.Time { return s.CreatedAt }
func (s *Task) GetUpdatedAt() time.Time { return s.UpdatedAt }

type List struct {
	Id   uint
	Name string
	// order of list in the board
	Order uint
}

type Board struct {
	Id    uint
	Tasks []*Task
	Lists []*List

	// settings contains:
	// - order of lists on frontend
	Settings string
}

type CalendarEntry struct {
	Id         uint
	Date       time.Time
	Name       string
	CalendarId uint
}

type Calendar struct {
	Id       uint
	Entries  []*CalendarEntry
	Settings string
}

type Document struct {
	Id       uint
	Name     string
	FolderId uint
}

type Folder struct {
	Id        uint
	Name      string
	Documents []*Document
	ParentId  *uint
	Children  []*Folder
}

type Project struct {
	Id   uint
	Name string

	Boards    []*Board
	Calendars []*Calendar
	Documents *Folder
}

type Comment struct {
	Id         uint
	Date       time.Time
	Content    string
	EntityType string
	EntityId   uint
	AuthorId   uint
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (s *Comment) GetId() uint                { return s.Id }
func (s *Comment) SetId(v uint)               { s.Id = v }
func (s *Comment) GetOwnerId() uint           { return s.AuthorId }
func (s *Comment) IsOwner(user AuthUser) bool { return user.GetId() == s.AuthorId }
func (s *Comment) GetCreatedAt() time.Time    { return s.CreatedAt }
func (s *Comment) GetUpdatedAt() time.Time    { return s.UpdatedAt }
