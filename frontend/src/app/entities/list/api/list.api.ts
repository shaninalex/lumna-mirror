import { inject, Injectable } from '@angular/core';
import { map, Observable } from 'rxjs';
import { APIResponse } from '@shared/models';
import { HttpClient } from '@angular/common/http';
import { ListModel, ListPayloadModel } from '@entities/list/model/list.model';

@Injectable({
    providedIn: 'root',
})
export class ListApi {
    http = inject(HttpClient);

    List(boardId: string): Observable<ListModel[]> {
        return this.http
            .get<
                APIResponse<ListModel[]>
            >(`/api/v1/board/${boardId}/lists`, { withCredentials: true })
            .pipe(map((response) => response.data));
    }

    Create(boardId: string, payload: ListPayloadModel): Observable<ListModel> {
        return this.http
            .post<
                APIResponse<ListModel>
            >(`/api/v1/board/${boardId}/lists`, payload, { withCredentials: true })
            .pipe(map((response) => response.data));
    }

    Delete(listId: string): Observable<void> {
        return this.http
            .delete<APIResponse<void>>(`/api/v1/lists/${listId}`, { withCredentials: true })
            .pipe(map((response) => response.data));
    }

    Patch(listId: string, payload: ListPayloadModel): Observable<ListModel> {
        return this.http
            .patch<
                APIResponse<ListModel>
            >(`/api/v1/lists/${listId}`, payload, { withCredentials: true })
            .pipe(map((response) => response.data));
    }
}
