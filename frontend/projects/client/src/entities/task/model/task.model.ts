export interface Task {
    id: number
    creator_id: number
    epic_id: number
    sprint_id: number
    project_id: number
    assignee: string
    completed: boolean
    title: string
    description: string
    status_id: number
    list_index: number
    code: string
    created_at: Date
    updated_at: Date
    deleted_at?: Date
}

export interface CreateTaskInput {
    title: string
    status_id: number
}
