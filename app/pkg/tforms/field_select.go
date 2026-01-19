package tforms

import (
	"fmt"
)

type SelectField struct {
	*BaseInput
	Options []*SelectableOption

	// Multiple in select should render something like "tags". But it required more JS
	Multiple bool
}

func NewSelectField(name string, options []*SelectableOption, multiple bool, required bool) *SelectField {
	s := &SelectField{
		BaseInput: &BaseInput{
			name:      name,
			inputType: InputTypeSelect,
			required:  required,
		},
		Multiple: multiple,
	}
	s.Options = s.copy(options)
	s.makeOptions(s.Options)
	return s
}

func (s *SelectField) Base() *BaseInput { return s.BaseInput }
func (s *SelectField) SetValue(v any) IBaseFormControl {
	for _, option := range s.Options {
		option.Selected = false
	}
	for _, option := range s.Options {
		if fmt.Sprintf("%v", option.Value) == fmt.Sprintf("%v", v) {
			option.Selected = true
			break
		}
	}
	s.makeOptions(s.Options)
	return s
}
func (s *SelectField) SetLabel(v string) IBaseFormControl {
	s.BaseInput.label = v
	return s
}
func (s *SelectField) SetPlaceholder(v string) IBaseFormControl {
	s.BaseInput.placeholder = v
	return s
}
func (s *SelectField) makeOptions(options []*SelectableOption) {
	strOptions := make([]Option, len(options))
	for i, option := range options {
		strOptions[i] = option.ToOption()
	}
	s.BaseInput.options = strOptions
}
func (s *SelectField) copy(options []*SelectableOption) []*SelectableOption {
	newOptions := make([]*SelectableOption, len(options))
	for i, o := range options {
		newOptions[i] = &SelectableOption{
			Label:    o.Label,
			Selected: o.Selected,
			Value:    o.Value,
		}
	}
	return newOptions
}
func (s *SelectField) SetError(e string) {
	s.BaseInput.inputError = e
}
