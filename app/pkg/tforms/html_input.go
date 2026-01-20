// IHTMLInput represent only html element. There are only getters.

package tforms

import (
	"fmt"
)

type IHTMLInput interface {
	Error() string
	HasError() bool
	HTMLInputType() string // TextInputType
	HTMLValue() string
	ID() string
	Label() string
	Name() string
	Options() []Option
	Placeholder() string
	Required() bool
	Type() InputType
}

type BaseInput struct {
	htmlType    TextInputType
	inputError  string
	inputType   InputType
	label       string
	name        string
	options     []Option
	placeholder string
	required    bool
	value       string
}

func (s *BaseInput) Error() string         { return s.inputError }
func (s *BaseInput) HasError() bool        { return len(s.inputError) > 0 }
func (s *BaseInput) HTMLInputType() string { return string(s.htmlType) }
func (s *BaseInput) HTMLValue() string     { return s.value }
func (s *BaseInput) ID() string            { return fmt.Sprintf("form_%s", s.Name()) }
func (s *BaseInput) Label() string         { return s.label }
func (s *BaseInput) Name() string          { return s.name }
func (s *BaseInput) Options() []Option     { return s.options }
func (s *BaseInput) Placeholder() string   { return s.placeholder }
func (s *BaseInput) Required() bool        { return s.required }
func (s *BaseInput) Type() InputType       { return s.inputType }
