export interface ProjectModel {
    id: number;
    name: string;
}

// Used to create/patch projects
export interface ProjectPayload {
    name: string;
}
