import { inject, Injectable } from "@angular/core";
import { map, Observable } from "rxjs";
import { APIResponse } from "@shared/models";
import { HttpClient } from "@angular/common/http";
import { TaskCreateModel, TaskListQueryModel, TaskModel, ToHttpParams } from "@entities/task";

@Injectable()
export class TaskApi {
    private http = inject(HttpClient);

    list(q: TaskListQueryModel): Observable<TaskModel[]> {
        const params = ToHttpParams(q);
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
