export interface BoardModel {
    id: string;
    title: string;
    order: number; // <= always positive
    projectID: string;
}

export interface BoardPayloadModel {
    title: string;
    projectID: string;
}
