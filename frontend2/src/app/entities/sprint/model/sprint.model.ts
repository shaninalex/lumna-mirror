import { HttpParams } from "@angular/common/http";

export interface SprintModel {
    id: number;
    title: string;
    body: string;
    done: boolean;
    project_id: number;
    started_at: Date;
    finished_at: Date;
    created_at: Date;
    updated_at: Date;
}
export interface SprintCreateModel {
    title: string;
    project_id: number;
}

export interface SprintListQueryModel {
    project_id?: number;
}

export function ToHttpParams(q: SprintListQueryModel): HttpParams {
    let params = new HttpParams();
    if (q.project_id) {
        params = params.set("project_id", q.project_id);
    }
    return params;
}
