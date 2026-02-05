export interface TaskModel {
    id: string;
    board_id: string;
    column_id: string;
    project_id: string;
    title: string;
    order: number;
    done: boolean;
    created_at: Date;
    updated_at: Date;
}

export interface TaskPayloadModel {
    title: string;
    order: number;
    column_id: string;
    project_id: string;
}
