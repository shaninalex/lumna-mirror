export interface Task {
	id: number
	user_id: number
	project_id: number
	status_id: number
	title: string
	completed: boolean
	description: string
	list_index: number
	code: string
	// badges:
	created_at: Date
	updated_at: Date
}

export interface CreateTaskInput {
	title: string
	status_id: number
}

export interface TaskDetailInput {
	title: string
	completed: boolean
	description: string
	list_index: number
	status_id: number
}
