export interface ProjectModel {
    id: number;
    title: string;
    slug: string;
    workspace_id: number;
}

export interface ProjectCreateModel {
    title: string;
    workspace_id: number;
}
