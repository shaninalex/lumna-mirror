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

    // This should be removed.
    // Activity have to be generated on backend side.
    activity?: {
        entity_type: string;
        entity_id: number;
        summary: string;
        action: string;
    };
}
