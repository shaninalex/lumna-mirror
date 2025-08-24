import {Issue} from '@client/entities/issue';

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
    issues: Issue[]
}
