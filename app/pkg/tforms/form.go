package tforms

import (
	"net/url"
)

type IForm interface {
	IsValid() bool
	SetValue(payload any)
	GetInputs() []IBaseFormControl
	GetActionUrl() string
	GetErrors() map[string]string
	SetErrors(e map[string]string)
	GetMethod() FormMethod
	SetMethod(v FormMethod)
	GetType() FormType
	SetType(v FormType)
}

type Form struct {
	actionUrl  string
	multipart  bool
	inputs     []IBaseFormControl
	formMethod FormMethod
	formType   FormType
}

func NewForm(action string, multipart bool, controls ...IBaseFormControl) IForm {
	form := &Form{
		actionUrl: action,
		multipart: multipart,
		inputs:    make([]IBaseFormControl, len(controls)),
	}

	for idx, input := range controls {
		form.inputs[idx] = input
	}

	return form
}

func (s *Form) IsValid() bool {
	for _, i := range s.inputs {
		if i.Base().HasError() {
			return false
		}
	}
	return true
}
func (s *Form) GetInputs() []IBaseFormControl { return s.inputs }
func (s *Form) GetActionUrl() string          { return s.actionUrl }
func (s *Form) SetMethod(v FormMethod)        { s.formMethod = v }
func (s *Form) SetType(v FormType)            { s.formType = v }

func (s *Form) GetErrors() map[string]string {
	errors := make(map[string]string, len(s.inputs))
	for _, input := range s.inputs {
		if input.Base().HasError() {
			errors[input.Base().Name()] = input.Base().inputError
		}
	}

	return errors
}

func (s *Form) GetMethod() FormMethod {
	if s.formMethod == "" {
		return FormMethodPost
	}
	return s.formMethod
}
func (s *Form) GetType() FormType {
	if s.formType == "" {
		return FormTypeXWWWFormUrlEncoded
	}
	return s.formType
}

func (s *Form) SetValue(payload any) {
	// request application/x-www-form-urlencoded form
	if values, ok := payload.(url.Values); ok {
		for k, v := range values {
			for _, input := range s.inputs {
				if input.Base().Name() == k {
					if input.Base().inputType == InputTypeCheckbox {
						input.SetValue(v)
					} else {
						input.SetValue(v[0])
					}
					break
				}
			}
		}
	}
}
func (s *Form) SetErrors(e map[string]string) {
	for fieldName, errorString := range e {
		for _, input := range s.inputs {
			if input.Base().Name() == fieldName {
				input.SetError(errorString)
				break
			}
		}
	}
}
