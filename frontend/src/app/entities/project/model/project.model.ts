export interface ProjectModel {
    id: number;
    title: string;
    created_at: Date;
    updated_at: Date;
}

// Used to create/patch projects
export interface ProjectPayload {
    title: string;
}
