import { inject, Injectable } from '@angular/core';
import type { Observable } from 'rxjs';
import { map } from 'rxjs';
import type { APIResponse } from '@shared/models';
import { HttpClient, HttpParams } from '@angular/common/http';
import type { BoardModel, BoardPayloadModel } from '../model/board.model';

@Injectable({
    providedIn: 'root',
})
export class BoardApi {
    http = inject(HttpClient);

    List(projectId: number): Observable<BoardModel[]> {
        let params = new HttpParams();
        params = params.set('project_id', projectId);
        return this.http
            .get<
                APIResponse<BoardModel[]>
            >(`/api/v1/lists`, { params: params, withCredentials: true })
            .pipe(map((response) => response.data));
    }

    Create(payload: BoardPayloadModel): Observable<BoardModel> {
        return this.http
            .post<APIResponse<BoardModel>>(`/api/v1/lists`, payload, { withCredentials: true })
            .pipe(map((response) => response.data));
    }

    Get(boardId: number): Observable<BoardModel> {
        return this.http
            .get<APIResponse<BoardModel>>(`/api/v1/lists/${boardId}`, { withCredentials: true })
            .pipe(map((response) => response.data));
    }

    Delete(boardId: number): Observable<void> {
        return this.http
            .delete<APIResponse<void>>(`/api/v1/lists/${boardId}`, { withCredentials: true })
            .pipe(map((response) => response.data));
    }

    Patch(boardId: number, payload: BoardPayloadModel): Observable<BoardModel> {
        return this.http
            .patch<
                APIResponse<BoardModel>
            >(`/api/v1/lists/${boardId}`, payload, { withCredentials: true })
            .pipe(map((response) => response.data));
    }
}
