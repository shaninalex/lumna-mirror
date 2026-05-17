import { inject, Injectable } from "@angular/core";
import { map, Observable } from "rxjs";
import { WorkspaceModel } from "../model/workspace.model";
import { APIResponse } from "@shared/models";
import { HttpClient } from "@angular/common/http";

@Injectable()
export class WorkspaceApi {
    private http = inject(HttpClient);

    list(): Observable<WorkspaceModel[]> {
        return this.http
            .get<
                APIResponse<WorkspaceModel[]>
            >(`/api/v1/workspaces`, { withCredentials: true })
            .pipe(map((response) => response.data));
    }
}
