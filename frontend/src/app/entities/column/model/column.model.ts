import { TaskModel } from '@entities/task';

export interface ColumnModel {
    id: number;
    title: string;
    order: number;
    board_id: number;
    project_id: number;
    created_at: Date;
    updated_at: Date;
}

export interface ColumnPayloadModel {
    title: string;
    order: number;
    board_id: number;
}
