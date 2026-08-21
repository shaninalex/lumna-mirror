import { inject, Injectable } from "@angular/core";
import type { Observable } from "rxjs";
import { map } from "rxjs";
import type { APIResponse } from "@shared/models";
import { HttpClient, HttpParams } from "@angular/common/http";
import type { StatusModel, StatusPayloadModel } from "../model/status.model";

@Injectable()
export class StatusApi {
    private http = inject(HttpClient);

    list(listId: number): Observable<StatusModel[]> {
        let params = new HttpParams()
        params = params.append('list_id', listId);
        return this.http
            .get<
                APIResponse<StatusModel[]>
            >(`/api/v1/statuses`, { params, withCredentials: true })
            .pipe(map((response) => response.data));
    }

    create(data: StatusPayloadModel): Observable<StatusModel> {
        return this.http
            .post<
                APIResponse<StatusModel>
            >(`/api/v1/statuses`, data, { withCredentials: true })
            .pipe(map((response) => response.data));
    }
}
