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

    activity?: {
        entity_type: string,
        entity_id: number,
        summary: string,
        action: string,
    }
}
