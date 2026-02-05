import { inject, Injectable } from '@angular/core';
import { map, Observable } from 'rxjs';
import { APIResponse } from '@shared/models';
import { HttpClient } from '@angular/common/http';
import { TaskModel, TaskPayloadModel } from '@entities/task';

@Injectable({
    providedIn: 'root',
})
export class TaskApi {
    http = inject(HttpClient);

    List(boardId: string): Observable<TaskModel[]> {
        return this.http
            .get<
                APIResponse<TaskModel[]>
            >(`/api/v1/board/${boardId}/tasks`, { withCredentials: true })
            .pipe(map((response) => response.data));
    }

    Create(payload: TaskPayloadModel): Observable<TaskModel> {
        return this.http
            .post<APIResponse<TaskModel>>(`/api/v1/tasks`, payload, { withCredentials: true })
            .pipe(map((response) => response.data));
    }

    Get(taskId: string): Observable<TaskModel> {
        return this.http
            .delete<APIResponse<TaskModel>>(`/api/v1/task/${taskId}`, { withCredentials: true })
            .pipe(map((response) => response.data));
    }

    Delete(taskId: string): Observable<void> {
        return this.http
            .delete<APIResponse<void>>(`/api/v1/task/${taskId}`, { withCredentials: true })
            .pipe(map((response) => response.data));
    }

    Patch(taskId: string, payload: TaskPayloadModel): Observable<TaskModel> {
        return this.http
            .patch<
                APIResponse<TaskModel>
            >(`/api/v1/task/${taskId}`, payload, { withCredentials: true })
            .pipe(map((response) => response.data));
    }
}
