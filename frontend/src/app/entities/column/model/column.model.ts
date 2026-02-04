import { TaskModel } from '@entities/task';

export interface ColumnModel {
    id: string;
    title: string;
    order: number;
    board_id: string;
    created_at: Date;
    updated_at: Date;
}

export interface ApiColumnModel {
    id: string;
    title: string;
    order: number;
    board_id: string;
    created_at: string;
    updated_at: string;
    tasks: TaskModel[];
}

export interface ColumnPayloadModel {
    title: string;
    order: number;
    board_id: string;
}
