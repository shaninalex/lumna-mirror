import { inject, Injectable } from "@angular/core";
import { map, Observable } from "rxjs";
import { APIResponse } from "@shared/models";
import { HttpClient } from "@angular/common/http";
import {
    SprintCreateModel,
    SprintListQueryModel,
    SprintModel,
    ToHttpParams
} from "@entities/sprint";

@Injectable()
export class SprintApi {
    private http = inject(HttpClient);

    list(q: SprintListQueryModel): Observable<SprintModel[]> {
        const params = ToHttpParams(q);
        return this.http
            .get<
                APIResponse<SprintModel[]>
            >(`/api/v1/sprints`, { params, withCredentials: true })
            .pipe(map((response) => response.data));
    }

    create(data: SprintCreateModel): Observable<SprintModel> {
        return this.http
            .post<
                APIResponse<SprintModel>
            >(`/api/v1/sprints`, data, { withCredentials: true })
            .pipe(map((response) => response.data));
    }
}
