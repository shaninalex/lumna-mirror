package tforms

type IBaseFormControl interface {
	Base() *BaseInput
	SetLabel(v string) IBaseFormControl
	SetPlaceholder(v string) IBaseFormControl
	SetValue(v any) IBaseFormControl
	SetError(e string)
}
