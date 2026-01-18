export interface KanbanBoardChangeOrderPayload {
    listId?: string;
    tasks?: Array<{
        id: string;
        order: number;
    }>;
    lists?: Array<{
        id: string;
        order?: number;
        tasks?: Array<{
            id: string;
            order: number;
        }>;
    }>;
}
