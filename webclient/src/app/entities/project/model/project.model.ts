export interface ProjectModel {
    id: number;
    name: string;
    created_at: Date;
    updated_at: Date;
}

// Used to create/patch projects
export interface ProjectPayload {
    name: string;
}
