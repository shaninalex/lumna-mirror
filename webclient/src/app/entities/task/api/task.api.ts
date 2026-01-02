import {inject, Injectable} from '@angular/core';
import {map, Observable} from 'rxjs';
import {APIResponse} from '@shared/models';
import {environment as env} from '@environments/environment.development';
import {HttpClient} from '@angular/common/http';
import {TaskModel, TaskPayloadModel} from '@entities/task';

@Injectable({
    providedIn: 'root',
})
export class TaskApi {
    http = inject(HttpClient)

    List(boardId: number): Observable<TaskModel[]> {
        return this.http.get<APIResponse<TaskModel[]>>(`${env.API_ROOT}/api/v1/board/${boardId}/tasks`, {withCredentials: true}).pipe(
            map(response => response.data)
        )
    }

    Create(boardId: number, payload: TaskPayloadModel): Observable<TaskModel> {
        return this.http.post<APIResponse<TaskModel>>(`${env.API_ROOT}/api/v1/task/${boardId}/tasks`, payload, {withCredentials: true}).pipe(
            map(response => response.data)
        )
    }

    Get(taskId: number): Observable<TaskModel> {
        return this.http.delete<APIResponse<TaskModel>>(`${env.API_ROOT}/api/v1/task/${taskId}`, {withCredentials: true}).pipe(
            map(response => response.data)
        )
    }

    Delete(taskId: number): Observable<void> {
        return this.http.delete<APIResponse<void>>(`${env.API_ROOT}/api/v1/task/${taskId}`, {withCredentials: true}).pipe(
            map(response => response.data)
        )
    }

    Patch(taskId: number, payload: TaskPayloadModel): Observable<TaskModel> {
        return this.http.patch<APIResponse<TaskModel>>(`${env.API_ROOT}/api/v1/task/${taskId}`, payload, {withCredentials: true}).pipe(
            map(response => response.data)
        )
    }
}
