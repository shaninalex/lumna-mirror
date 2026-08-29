export interface ColumnModel {
    id: number;
    title: number;
    board_id: number;
    meta: ColumnMeta;
    created_at: Date;
    update_at: Date;
}

export interface ColumnPayloadModel {
    title: string;
    order: number;
    board_id: number;
}

export interface ColumnMeta {
    order: number;
    color: string;
    icon: string;
    expanded: boolean;
}
