export interface TaskModel {
    id: number
    board_id: number
    list_id: number
    name: string
    order: number
    done: boolean
    created_at: Date
    updated_at: Date
}

export interface TaskPayloadModel {
    name: string
    list_id: number
}
