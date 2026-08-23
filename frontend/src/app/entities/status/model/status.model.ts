export interface StatusModel {
    id: number;
    title: number;
    project_id: number;
    list_id: number;
    meta: StatusMeta;
    created_at: Date;
    update_at: Date;
}

export interface StatusPayloadModel {
    title: string;
    order: number;
    project_id: number;
    list_id: number;
}

export interface StatusMeta {
    order: number;
    color: string;
    icon: string;
    expanded: boolean;
}
