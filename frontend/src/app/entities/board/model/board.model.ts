export interface BoardModel {
    id: number;
    title: string;
    project_id: number;

    created_at: Date;
    updated_at: Date;
}

export interface BoardPayloadModel {
    title: string;
    project_id: number;
}
