export interface TaskModel {
    id: number;
    column_id: number;
    project_id: number;
    title: string;
    order: number;
    done: boolean;
    created_at: Date;
    updated_at: Date;
}

export interface TaskPayloadModel {
    title: string;
    order: number;
    column_id: number;
}
