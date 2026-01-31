import { inject, Injectable } from '@angular/core';
import { map, Observable } from 'rxjs';
import { APIResponse } from '@shared/models';
import { HttpClient } from '@angular/common/http';
import { ColumnModel, ColumnPayloadModel } from '../model/column.model';

@Injectable({
    providedIn: 'root',
})
export class ColumnsApi {
    http = inject(HttpClient);

    List(boardId: string): Observable<ColumnModel[]> {
        return this.http
            .get<
                APIResponse<ColumnModel[]>
            >(`/api/v1/board/${boardId}/columns`, { withCredentials: true })
            .pipe(map((response) => response.data));
    }

    Create(boardId: string, payload: ColumnPayloadModel): Observable<ColumnModel> {
        return this.http
            .post<
                APIResponse<ColumnModel>
            >(`/api/v1/board/${boardId}/columns`, payload, { withCredentials: true })
            .pipe(map((response) => response.data));
    }

    Delete(columnId: string): Observable<void> {
        return this.http
            .delete<APIResponse<void>>(`/api/v1/columns/${columnId}`, { withCredentials: true })
            .pipe(map((response) => response.data));
    }

    Patch(columnId: string, payload: ColumnPayloadModel): Observable<ColumnModel> {
        return this.http
            .patch<
                APIResponse<ColumnModel>
            >(`/api/v1/columns/${columnId}`, payload, { withCredentials: true })
            .pipe(map((response) => response.data));
    }
}
