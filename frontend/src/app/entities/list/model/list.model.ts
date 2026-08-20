export interface ListModel {
    id: number;
    title: string;
    project_id: number;

    created_at: Date;
    updated_at: Date;
}

export interface ListPayloadModel {
    title: string;
    project_id: number;
}
