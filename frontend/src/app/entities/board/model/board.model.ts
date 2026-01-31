export interface BoardModel {
    id: string;
    title: string;
    project_id: string;

    created_at: Date;
    updated_at: Date;
}

export interface BoardPayloadModel {
    title: string;
    projectID: string;
}
