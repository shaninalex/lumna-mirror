package app

type TaskFilter struct {
	Project string `query:"project,required"`
}
