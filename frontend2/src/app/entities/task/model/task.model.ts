import { HttpParams } from "@angular/common/http";

export interface TaskModel {
    id: number;
    title: string;
    order: number;
    done: boolean;
    body: string;

    status_id: number;
    project_id: number;

    created_at: Date;
    updated_at: Date;
}

export interface TaskCreateModel {
    title: string;
    project_id: number;
    order?: number;
    status_id?: number;
}

export interface TaskListQueryModel {
    project_id?: number;
}

export function ToHttpParams(q: TaskListQueryModel): HttpParams {
    let params = new HttpParams();
    if (q.project_id) {
        params = params.set("project_id", q.project_id);
    }
    return params;
}
