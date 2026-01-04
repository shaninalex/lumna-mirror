import { inject, Injectable } from '@angular/core';
import { map, Observable, tap } from 'rxjs';
import { APIResponse } from '@shared/models';
import { environment as env } from '@environments/environment.development';
import { HttpClient } from '@angular/common/http';
import { KanbanBoardChangeOrderPayload } from '../model/kanban.model';


@Injectable()
export class KanbanApi {
    http = inject(HttpClient)

    public Patch(boardId: number, payload: KanbanBoardChangeOrderPayload): Observable<any> {
        return this.http.patch<APIResponse<any>>(`${env.API_ROOT}/api/v1/board/${boardId}/order`, payload, {withCredentials: true}).pipe(
            tap(data => console.log(data)),
            map(response => response.data)
        )
    }
}
