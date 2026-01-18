export interface ProjectModel {
    id: string;
    title: string;
    created_at: Date;
    updated_at: Date;
}

// Used to create/patch projects
export interface ProjectPayload {
    title: string;
}
