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
    http = inject(HttpClient)

    List(projectId: number): Observable<BoardModel[]> {
        return this.http.get<APIResponse<BoardModel[]>>(`${env.API_ROOT}/api/v1/project/${projectId}/boards`, { withCredentials: true }).pipe(
            map(response => response.data)
        )
    }

    Create(projectId: number, payload: BoardPayloadModel): Observable<BoardModel[]> {
        return this.http.post<APIResponse<BoardModel[]>>(`${env.API_ROOT}/api/v1/project/${projectId}/boards`, payload, { withCredentials: true }).pipe(
            map(response => response.data)
        )
    }

    DeleteProject(boardId: number): Observable<void> {
        return this.http.delete<APIResponse<void>>(`${env.API_ROOT}/api/v1/board/${boardId}`, {withCredentials: true}).pipe(
            map(response => response.data)
        )
    }

    Patch(boardId: number, payload: BoardPayloadModel): Observable<BoardModel> {
        return this.http.patch<APIResponse<BoardModel>>(`${env.API_ROOT}/api/v1/board/${boardId}`, payload, {withCredentials: true}).pipe(
            map(response => response.data)
        )
    }
}
