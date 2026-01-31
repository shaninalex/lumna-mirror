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
