import { HttpClient } from '@angular/common/http';
import { Injectable, inject } from '@angular/core';
import type { ColumnModel } from '@entities/column';
import type { APIResponse } from '@shared/models';
import { map, type Observable } from 'rxjs';

@Injectable()
export class KanbanApi {
    private http = inject(HttpClient);

    SetColumns(data: unknown): Observable<ColumnModel[]> {
        return this.http
            .post<APIResponse<ColumnModel[]>>(`/api/v1/columns/reorder`, data, { withCredentials: true })
            .pipe(map((response) => response.data));
    }
}
