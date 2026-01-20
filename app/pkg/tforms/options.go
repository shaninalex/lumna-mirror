package tforms

import "fmt"

type Option struct {
	Value    string
	Label    string
	Selected bool
}

type SelectableOption struct {
	Label    string
	Value    any
	Selected bool
}

func (s *SelectableOption) ToOption() Option {
	return Option{
		Label:    s.Label,
		Selected: s.Selected,
		Value:    fmt.Sprintf("%v", s.Value),
	}
}
