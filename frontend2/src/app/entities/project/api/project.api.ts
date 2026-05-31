import { inject, Injectable } from "@angular/core";
import { map, Observable } from "rxjs";
import { APIResponse } from "@shared/models";
import { HttpClient } from "@angular/common/http";
import { ProjectCreateModel, ProjectModel } from "@entities/project";

@Injectable()
export class ProjectApi {
    private http = inject(HttpClient);

    list(): Observable<ProjectModel[]> {
        return this.http
            .get<
                APIResponse<ProjectModel[]>
            >(`/api/v1/projects`, { withCredentials: true })
            .pipe(map((response) => response.data));
    }

    create(data: ProjectCreateModel): Observable<ProjectModel> {
        return this.http
            .post<
                APIResponse<ProjectModel>
            >(`/api/v1/projects`, data, { withCredentials: true })
            .pipe(map((response) => response.data));
    }
}
