import { inject, Injectable } from '@angular/core';
import { map, Observable } from 'rxjs';
import { APIResponse } from '@shared/models';
import { HttpClient, HttpParams } from '@angular/common/http';
import { ApiColumnModel, ColumnModel, ColumnPayloadModel } from '../model/column.model';

@Injectable({
    providedIn: 'root',
})
export class ColumnsApi {
    http = inject(HttpClient);

    List(boardId: number): Observable<ApiColumnModel[]> {
        const params = new HttpParams().set('board_id', boardId);
        return this.http
            .get<
                APIResponse<ApiColumnModel[]>
            >(`/api/v1/columns`, { params, withCredentials: true })
            .pipe(map((response) => response.data));
    }

    Create(payload: ColumnPayloadModel): Observable<ColumnModel> {
        return this.http
            .post<APIResponse<ColumnModel>>(`/api/v1/columns`, payload, { withCredentials: true })
            .pipe(map((response) => response.data));
    }

    Delete(columnId: number): Observable<void> {
        return this.http
            .delete<APIResponse<void>>(`/api/v1/column/${columnId}`, { withCredentials: true })
            .pipe(map((response) => response.data));
    }

    Patch(columnId: number, payload: ColumnPayloadModel): Observable<ColumnModel> {
        return this.http
            .patch<
                APIResponse<ColumnModel>
            >(`/api/v1/column/${columnId}`, payload, { withCredentials: true })
            .pipe(map((response) => response.data));
    }
}
