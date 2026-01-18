export interface ListModel {
    id: string;
    title: string;
    order: number;
    board_id: string;
}

export interface ListPayloadModel {
    title: string;
    order: number;
}
