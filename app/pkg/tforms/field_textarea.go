package tforms

import (
	"fmt"
)

// TextArea represents a multi-line text input.
type TextArea struct {
	*BaseInput
}

func NewTextArea(name string, required bool) *TextArea {
	return &TextArea{
		BaseInput: &BaseInput{
			name:      name,
			inputType: InputTypeTextArea,
			required:  required,
		},
	}
}

func (s *TextArea) Base() *BaseInput { return s.BaseInput }
func (s *TextArea) SetValue(v any) IBaseFormControl {
	s.BaseInput.value = fmt.Sprintf("%v", v)
	return s
}
func (s *TextArea) SetLabel(v string) IBaseFormControl {
	s.BaseInput.label = v
	return s
}
func (s *TextArea) SetPlaceholder(v string) IBaseFormControl {
	s.BaseInput.placeholder = v
	return s
}
func (s *TextArea) SetError(e string) {
	s.BaseInput.inputError = e
}
