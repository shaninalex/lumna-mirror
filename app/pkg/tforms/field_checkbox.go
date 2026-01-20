package tforms

import (
	"fmt"
)

// CheckboxField represents a checkbox field.
type CheckboxField struct {
	*BaseInput
	Options []*SelectableOption

	// If false it will render radiobuttons instead
	Multiple bool
}

func NewSelectableField(name string, selectableType InputType, options []*SelectableOption, multiple bool, required bool) *CheckboxField {
	s := &CheckboxField{
		BaseInput: &BaseInput{
			name:      name,
			inputType: selectableType,
			required:  required,
		},
		Multiple: multiple,
	}
	s.Options = s.copy(options)
	s.makeOptions(options)
	return s
}

func (s *CheckboxField) Base() *BaseInput { return s.BaseInput }
func (s *CheckboxField) SetValue(v any) IBaseFormControl {
	for _, option := range s.Options {
		option.Selected = false
	}
	if values, ok := v.([]string); ok {
		for _, value := range values {
			for _, option := range s.Options {
				if fmt.Sprintf("%v", option.Value) == fmt.Sprintf("%v", value) {
					option.Selected = true
					break
				}
			}
		}
		s.makeOptions(s.Options)
		return s
	}
	return s
}
func (s *CheckboxField) SetLabel(v string) IBaseFormControl {
	s.BaseInput.label = v
	return s
}
func (s *CheckboxField) SetPlaceholder(v string) IBaseFormControl {
	s.BaseInput.placeholder = v
	return s
}

// private

func (s *CheckboxField) makeOptions(options []*SelectableOption) {
	strOptions := make([]Option, len(options))
	for i, option := range options {
		strOptions[i] = option.ToOption()
	}
	s.BaseInput.options = strOptions
}
func (s *CheckboxField) copy(options []*SelectableOption) []*SelectableOption {
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
func (s *CheckboxField) SetError(e string) {
	s.BaseInput.inputError = e
}
