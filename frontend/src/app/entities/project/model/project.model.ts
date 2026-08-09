export interface ProjectModel {
    id: number
    title: string
    key: string
    workspace_id: number
    owner_id: number
    meta: string

    created_at: Date;
    updated_at: Date;
}

// Used to create/patch projects
export interface ProjectCreateModel {
    title: string;
    workspace_id: number;
}
