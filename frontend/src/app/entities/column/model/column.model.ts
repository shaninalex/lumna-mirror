import { TaskModel } from '@entities/task';

export interface ColumnModel {
    id: string;
    title: string;
    order: number;
    board_id: string;
    tasks: TaskModel[];
    created_at: Date;
    updated_at: Date;
}

export interface ColumnPayloadModel {
    title: string;
    order: number;
}
