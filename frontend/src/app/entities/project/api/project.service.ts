import { inject, Injectable } from '@angular/core';
import { map, Observable } from 'rxjs';
import { ProjectModel, ProjectCreateModel } from '../model/project.model';
import { APIResponse } from '@shared/models';
import { HttpClient, HttpParams } from '@angular/common/http';

@Injectable({
    providedIn: 'root',
})
export class ProjectApi {
    http = inject(HttpClient);

    GetProjects(workspaceId: number): Observable<ProjectModel[]> {
        return this.http
            .get<APIResponse<ProjectModel[]>>(`/api/v1/projects`, {
                params: new HttpParams().set('workspace_id', workspaceId),
                withCredentials: true,
            })
            .pipe(map((response) => response.data));
    }

    GetProject(projectId: number): Observable<ProjectModel> {
        return this.http
            .get<APIResponse<ProjectModel>>(`/api/v1/projects/${projectId}`, { withCredentials: true })
            .pipe(map((response) => response.data));
    }

    CreateProject(payload: ProjectCreateModel): Observable<ProjectModel> {
        return this.http
            .post<APIResponse<ProjectModel>>(`/api/v1/projects`, payload, { withCredentials: true })
            .pipe(map((response) => response.data));
    }

    DeleteProject(projectId: number): Observable<void> {
        return this.http
            .delete<APIResponse<void>>(`/api/v1/projects/${projectId}`, { withCredentials: true })
            .pipe(map((response) => response.data));
    }

    Patch(projectId: number, payload: ProjectCreateModel): Observable<ProjectModel> {
        return this.http
            .patch<
                APIResponse<ProjectModel>
            >(`/api/v1/projects/${projectId}`, payload, { withCredentials: true })
            .pipe(map((response) => response.data));
    }
}
