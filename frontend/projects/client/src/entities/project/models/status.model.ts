import {Task} from '@client/entities/task';

export interface StatusConfig {
    color: string
}

export interface Status {
    complete: boolean
    config: StatusConfig
    description: string
    id: string
    index: number
    title: string
    tasks: Task[]
}
