import { TaskModel } from '@entities/task';

export interface ColumnModel {
    id: number;
    title: string;
    order: number;
    board_id: number;
    created_at: Date;
    updated_at: Date;
}

export interface ApiColumnModel {
    id: number;
    title: string;
    order: number;
    board_id: number;
    created_at: Date;
    updated_at: Date;
    tasks: TaskModel[];
}

export interface ColumnPayloadModel {
    title: string;
    order: number;
    board_id: number;
}
