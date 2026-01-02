import { inject, Injectable } from '@angular/core';
import { map, Observable } from 'rxjs';
import { APIResponse } from '@shared/models';
import { environment as env } from '@environments/environment.development';
import { HttpClient } from '@angular/common/http';

/*
DOCS: https://gist.github.com/shaninalex/37aa0ad6eb8d3848d05b8ea7d1e72389
*/

@Injectable()
export class KanbanApi {
    http = inject(HttpClient)

    Patch(boardId: number, payload: any): Observable<any> {
        return this.http.patch<APIResponse<any>>(`${env.API_ROOT}/api/v1/board/${boardId}/kanban`, payload, {withCredentials: true}).pipe(
            map(response => response.data)
        )
    }
}
