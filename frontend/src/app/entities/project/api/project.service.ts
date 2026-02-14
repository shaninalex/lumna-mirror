import { inject, Injectable } from '@angular/core';
import { map, Observable } from 'rxjs';
import { ProjectModel, ProjectPayload } from '@entities/project';
import { APIResponse } from '@shared/models';
import { HttpClient } from '@angular/common/http';

@Injectable({
    providedIn: 'root',
})
export class ProjectApi {
    http = inject(HttpClient);

    GetProjects(): Observable<ProjectModel[]> {
        return this.http
            .get<APIResponse<ProjectModel[]>>(`/api/v1/projects`, { withCredentials: true })
            .pipe(map((response) => response.data));
    }

    CreateProject(payload: ProjectPayload): Observable<ProjectModel> {
        return this.http
            .post<APIResponse<ProjectModel>>(`/api/v1/projects`, payload, { withCredentials: true })
            .pipe(map((response) => response.data));
    }

    DeleteProject(projectId: string): Observable<void> {
        return this.http
            .delete<APIResponse<void>>(`/api/v1/project/${projectId}`, { withCredentials: true })
            .pipe(map((response) => response.data));
    }

    Patch(projectId: string, payload: ProjectPayload): Observable<ProjectModel> {
        return this.http
            .patch<
                APIResponse<ProjectModel>
            >(`/api/v1/project/${projectId}`, payload, { withCredentials: true })
            .pipe(map((response) => response.data));
    }
}
