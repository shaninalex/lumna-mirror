export interface StatusModel {
    id: number
    title: number
    order: number
    project_id: number
    list_id: number
    created_at: Date
    update_at: Date
}

export interface StatusPayloadModel {
    title: number
    order: number
    project_id: number
    list_id: number
}