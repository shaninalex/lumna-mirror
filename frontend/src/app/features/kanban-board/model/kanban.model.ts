export interface KanbanBoardChangeOrderPayload {
    columnId?: number;
    moveType: string;

    tasks?: Array<{
        id: number;
        order: number;
    }>;

    columns?: Array<{
        id: number;
        order?: number;
        tasks?: Array<{
            id: number;
            order: number;
        }>;
    }>;
}
