export interface KanbanBoardChangeOrderPayload {
    listId?: number
    tasks?: Array<{
        id: number,
        order: number
    }>
    lists?: Array<{
        id: number,
        order?: number,
        tasks?: Array<{
            id: number,
            order: number
        }>
    }>
}
