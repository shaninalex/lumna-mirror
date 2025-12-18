package models

import "time"

type CalendarEntry struct {
	Id         uint
	Date       time.Time
	Name       string
	CalendarId uint
}

func (s *CalendarEntry) GetId() uint  { return s.Id }
func (s *CalendarEntry) SetId(u uint) { s.Id = u }

type Calendar struct {
	Id       uint
	Entries  []*CalendarEntry
	Settings string
}

func (s *Calendar) GetId() uint  { return s.Id }
func (s *Calendar) SetId(u uint) { s.Id = u }

