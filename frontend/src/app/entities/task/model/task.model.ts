export interface TaskModel {
    id: string;
    board_id: string;
    column_id: string;
    title: string;
    order: number;
    done: boolean;
    created_at: Date;
    updated_at: Date;
}

export interface TaskPayloadModel {
    title: string;
    column_id: string;
    order: number;
}

export interface KanbanBoardChangeOrderPayload {
    columnId?: string;
    tasks?: Array<{
        id: string;
        order: number;
    }>;
    columns?: Array<{
        id: string;
        order?: number;
        tasks?: Array<{
            id: string;
            order: number;
        }>;
    }>;
}
