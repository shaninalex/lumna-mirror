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
    created_at: Date
    updated_at: Date

	comments: Comment[]
	comments_count: number
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

export interface Comment {
	id: number
	task_id: number
	user_id: number
	content: string
	created_at: Date
}
