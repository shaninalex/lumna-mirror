package tforms

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation"
)

// InputField represents text-based input fields.
type InputField struct {
	*BaseInput
	Rules []validation.Rule
}

// NewInputField creates a new text-based input field.
func NewInputField(name string, textInputType TextInputType, required bool) *InputField {
	return &InputField{
		BaseInput: &BaseInput{
			name:      name,
			inputType: InputTypeText,
			htmlType:  textInputType,
			required:  required,
		},
	}
}

// public

func (s *InputField) Base() *BaseInput { return s.BaseInput }
func (s *InputField) SetValue(v any) IBaseFormControl {
	s.BaseInput.value = fmt.Sprintf("%v", v)
	return s
}
func (s *InputField) SetLabel(v string) IBaseFormControl {
	s.BaseInput.label = v
	return s
}
func (s *InputField) SetPlaceholder(v string) IBaseFormControl {
	s.BaseInput.placeholder = v
	return s
}
func (s *InputField) SetError(e string) {
	s.BaseInput.inputError = e
}
