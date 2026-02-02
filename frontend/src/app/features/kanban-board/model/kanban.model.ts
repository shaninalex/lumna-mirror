export interface KanbanBoardChangeOrderPayload {
    columnId?: string;
    moveType: string;

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
