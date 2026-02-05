import { inject, Injectable } from '@angular/core';
import { map, Observable } from 'rxjs';
import { APIResponse } from '@shared/models';
import { environment as env } from '@environments/environment.development';
import { HttpClient } from '@angular/common/http';
import { BoardModel, BoardPayloadModel } from '../model/board.model';

@Injectable({
    providedIn: 'root',
})
export class BoardApi {
    http = inject(HttpClient);

    List(projectId: string): Observable<BoardModel[]> {
        return this.http
            .get<
                APIResponse<BoardModel[]>
            >(`/api/v1/project/${projectId}/boards`, { withCredentials: true })
            .pipe(map((response) => response.data));
    }

    Create(payload: BoardPayloadModel): Observable<BoardModel> {
        return this.http
            .post<
                APIResponse<BoardModel>
            >(`/api/v1/boards`, payload, { withCredentials: true })
            .pipe(map((response) => response.data));
    }

    Delete(boardId: string): Observable<void> {
        return this.http
            .delete<APIResponse<void>>(`/api/v1/board/${boardId}`, { withCredentials: true })
            .pipe(map((response) => response.data));
    }

    Patch(boardId: string, payload: BoardPayloadModel): Observable<BoardModel> {
        return this.http
            .patch<
                APIResponse<BoardModel>
            >(`/api/v1/board/${boardId}`, payload, { withCredentials: true })
            .pipe(map((response) => response.data));
    }
}
