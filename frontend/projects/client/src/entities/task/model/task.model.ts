import {v4 as uuidv4} from 'uuid';

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
    list_idx: number
    code: string
    created_at: Date
    updated_at: Date
    deleted_at?: Date
}

function NewTask(): Task {
    return {
        id: uuidv4(),
        creator_id: "",
        epic_id: "",
        sprint_id: "",
        project_id: "",
        assignee: "",
        completed: false,
        title: "",
        description: "",
        status: "",
        list_idx: 0,
        code: `task-${Math.floor(100000 + Math.random() * 900000)}`,
        created_at: new Date(),
        updated_at: new Date(),
    }
}
