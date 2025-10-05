export interface Status {
    id: number
    title: string
    complete: boolean
    project_id: number
    index: number
    config: StatusConfig
}

export interface StatusConfig {
    color: string
}

export interface StatusInput {
    title: string
    complete: boolean
}
