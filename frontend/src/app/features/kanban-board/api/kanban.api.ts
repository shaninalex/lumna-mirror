import { inject, Injectable } from '@angular/core';
import { map, Observable } from 'rxjs';
import { APIResponse } from '@shared/models';
import { HttpClient } from '@angular/common/http';
import { KanbanBoardChangeOrderPayload } from '../model/kanban.model';

@Injectable()
export class KanbanApi {
    http = inject(HttpClient);

    public Patch(boardId: number, payload: KanbanBoardChangeOrderPayload): Observable<any> {
        return this.http
            .patch<
                APIResponse<any>
            >(`/api/v1/board/${boardId}/order`, payload, { withCredentials: true })
            .pipe(map((response) => response.data));
    }
}
