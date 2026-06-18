export interface ProjectModel {
    id: number;
    title: string;
    workspace_id: number;

    appLink: string;
}

export interface ProjectCreateModel {
    title: string;
    workspace_id: number;
}
