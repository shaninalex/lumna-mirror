import { inject, Injectable } from '@angular/core';
import { map, Observable } from 'rxjs';
import { APIResponse } from '@shared/models';
import { environment as env } from '@environments/environment.development';
import { HttpClient } from '@angular/common/http';
import {ListModel, ListPayloadModel} from '@entities/list/model/list.model';

@Injectable({
    providedIn: 'root',
})
export class ListApi {
    http = inject(HttpClient)

    List(boardId: number): Observable<ListModel[]> {
        return this.http.get<APIResponse<ListModel[]>>(`${env.API_ROOT}/api/v1/board/${boardId}/lists`, { withCredentials: true }).pipe(
            map(response => response.data)
        )
    }

    Create(boardId: number, payload: ListPayloadModel): Observable<ListModel> {
        return this.http.post<APIResponse<ListModel>>(`${env.API_ROOT}/api/v1/board/${boardId}/lists`, payload, { withCredentials: true }).pipe(
            map(response => response.data)
        )
    }

    Delete(listId: number): Observable<void> {
        return this.http.delete<APIResponse<void>>(`${env.API_ROOT}/api/v1/lists/${listId}`, {withCredentials: true}).pipe(
            map(response => response.data)
        )
    }

    Patch(listId: number, payload: ListPayloadModel): Observable<ListModel> {
        return this.http.patch<APIResponse<ListModel>>(`${env.API_ROOT}/api/v1/lists/${listId}`, payload, {withCredentials: true}).pipe(
            map(response => response.data)
        )
    }
}
