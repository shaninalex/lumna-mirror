export interface Task {
    id: number
    creator_id: string
    epic_id: string
    sprint_id: string
    project_id: string
    assignee: string
    completed: boolean
    title: string
    description: string
    status: string
    list_idx: number
    code: string
    created_at: Date
    updated_at: Date
    deleted_at?: Date
}

export interface CreateTaskDto {
    title: string
    status_id: string
    project_code: string
}
