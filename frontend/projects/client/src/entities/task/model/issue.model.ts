export interface Task {
    id: string
    creator_id: string
    epic_id: string
    sprint_id: string
    project_id: string
    assignee: string
    completed: boolean
    title: string
    description: string
    status: string
    created_at: Date
    updated_at: Date
    deleted_at?: Date
}
