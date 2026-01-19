package tforms

import (
	"fmt"
)

type HiddenField struct {
	*BaseInput
}

func NewHiddenField(name, value string, required bool) *HiddenField {
	return &HiddenField{
		BaseInput: &BaseInput{
			name:      name,
			inputType: InputTypeHidden,
			required:  required,
			value:     value,
		},
	}
}

func (s *HiddenField) Base() *BaseInput { return s.BaseInput }
func (s *HiddenField) SetValue(v any) IBaseFormControl {
	s.BaseInput.value = fmt.Sprintf("%v", v)
	return s
}
func (s *HiddenField) SetLabel(v string) IBaseFormControl       { return s }
func (s *HiddenField) SetPlaceholder(v string) IBaseFormControl { return s }
func (s *HiddenField) SetError(e string) {
	s.BaseInput.inputError = e
}
