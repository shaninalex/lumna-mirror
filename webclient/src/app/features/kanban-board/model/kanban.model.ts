export interface KanbanBoardChangeOrderPayload {
    lists: Array<{
        id: number,
        order?: number,
        tasks?: Array<{
            id: number,
            order: number
        }>
    }>
}
