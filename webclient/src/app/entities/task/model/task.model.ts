export interface TaskModel {
    id: number
    name: string
    list_id: number
    done: boolean
    order: number
    created_at: Date
    updated_at: Date
}

export interface TaskPayloadModel {
    name: string
    list_id: number
}
