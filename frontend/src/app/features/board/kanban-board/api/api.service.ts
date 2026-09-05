import { HttpClient } from '@angular/common/http';
import { Injectable, inject } from '@angular/core';
import { ColumnModel } from '@entities/column';
import { TaskModel } from '@entities/task';
import { APIResponse } from '@shared/models';
import { map, Observable } from 'rxjs';

@Injectable()
export class KanbanApi {
    private http = inject(HttpClient);

    SetColumns(data: unknown): Observable<ColumnModel[]> {
        return this.http
            .post<APIResponse<ColumnModel[]>>(`/api/v1/columns/reorder`, data, { withCredentials: true })
            .pipe(map((response) => response.data));
    }

    MoveTask(data: unknown): Observable<TaskModel[]> {
        return this.http
            .post<APIResponse<TaskModel[]>>(`/api/v1/tasks/move`, data, { withCredentials: true })
            .pipe(map((response) => response.data));
    }

    TransferTask(data: unknown): Observable<TaskModel[]> {
        return this.http
            .post<APIResponse<TaskModel[]>>(`/api/v1/tasks/transfer`, data, { withCredentials: true })
            .pipe(map((response) => response.data));
    }
}
