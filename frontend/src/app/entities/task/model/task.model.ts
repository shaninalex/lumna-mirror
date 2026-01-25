export interface TaskModel {
    id: string;
    board_id: string;
    list_id: string;
    title: string;
    order: number;
    done: boolean;
    created_at: Date;
    updated_at: Date;
}

export interface TaskPayloadModel {
    title: string;
    list_id: string;
    order: number;
}
