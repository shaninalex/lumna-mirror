import { inject, Injectable } from "@angular/core";
import { map, Observable } from "rxjs";
import { WorkspaceCreateModel, WorkspaceModel } from "../model/workspace.model";
import { APIResponse } from "@shared/models";
import { HttpClient, HttpParams } from "@angular/common/http";

@Injectable()
export class WorkspaceApi {
    private http = inject(HttpClient);

    list(active: boolean): Observable<WorkspaceModel[]> {
        let params = new HttpParams();
        params = params.set("active", active);
        return this.http
            .get<
                APIResponse<WorkspaceModel[]>
            >(`/api/v1/workspaces`, { params, withCredentials: true })
            .pipe(map((response) => response.data));
    }

    create(data: WorkspaceCreateModel): Observable<WorkspaceModel> {
        return this.http
            .post<
                APIResponse<WorkspaceModel>
            >(`/api/v1/workspaces`, data, { withCredentials: true })
            .pipe(map((response) => response.data));
    }
}
