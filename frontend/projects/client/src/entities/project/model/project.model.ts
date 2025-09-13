export interface Project {
    id: string
    title: string
    project_key: string
    statuses: TaskStatus[]
    created_at: Date
    updated_at: Date
}

export interface TaskStatus {
    id: string
    title: string
    description: string
    complete :boolean
    index: number
    config: TaskStatusConfig
}

export interface TaskStatusConfig {
    color: string
}
