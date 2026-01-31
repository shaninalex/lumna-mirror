export interface ColumnModel {
    id: string;
    title: string;
    order: number;
    board_id: string;
}

export interface ColumnPayloadModel {
    title: string;
    order: number;
}
