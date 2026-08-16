import { inject, Injectable } from "@angular/core";
import type { Observable } from "rxjs";
import { map } from "rxjs";
import type { APIResponse } from "@shared/models";
import { HttpClient } from "@angular/common/http";
import type { TaskCreateModel, TaskListQueryModel, TaskModel } from "../model/task.model";
import { toHttpParams } from "../model/task.model";

@Injectable()
export class TaskApi {
    private http = inject(HttpClient);

    list(q: TaskListQueryModel): Observable<TaskModel[]> {
        const params = toHttpParams(q);
        return this.http
            .get<
                APIResponse<TaskModel[]>
            >(`/api/v1/tasks`, { params, withCredentials: true })
            .pipe(map((response) => response.data));
    }

    create(data: TaskCreateModel): Observable<TaskModel> {
        return this.http
            .post<
                APIResponse<TaskModel>
            >(`/api/v1/tasks`, data, { withCredentials: true })
            .pipe(map((response) => response.data));
    }
}
