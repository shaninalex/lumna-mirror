import { inject, Injectable } from '@angular/core';
import type { Observable } from 'rxjs';
import { map } from 'rxjs';
import type { APIResponse } from '@shared/models';
import { HttpClient, HttpParams } from '@angular/common/http';
import type { ListModel, ListPayloadModel } from '../model/list.model';

@Injectable({
    providedIn: 'root',
})
export class ListApi {
    http = inject(HttpClient);

    List(projectId: number): Observable<ListModel[]> {
        let params = new HttpParams();
        params = params.set('project_id', projectId);
        return this.http
            .get<
                APIResponse<ListModel[]>
            >(`/api/v1/lists`, { params: params, withCredentials: true })
            .pipe(map((response) => response.data));
    }

    Create(payload: ListPayloadModel): Observable<ListModel> {
        return this.http
            .post<APIResponse<ListModel>>(`/api/v1/lists`, payload, { withCredentials: true })
            .pipe(map((response) => response.data));
    }

    Get(boardId: number): Observable<ListModel> {
        return this.http
            .get<APIResponse<ListModel>>(`/api/v1/lists/${boardId}`, { withCredentials: true })
            .pipe(map((response) => response.data));
    }

    Delete(boardId: number): Observable<void> {
        return this.http
            .delete<APIResponse<void>>(`/api/v1/lists/${boardId}`, { withCredentials: true })
            .pipe(map((response) => response.data));
    }

    Patch(boardId: number, payload: ListPayloadModel): Observable<ListModel> {
        return this.http
            .patch<
                APIResponse<ListModel>
            >(`/api/v1/lists/${boardId}`, payload, { withCredentials: true })
            .pipe(map((response) => response.data));
    }
}
