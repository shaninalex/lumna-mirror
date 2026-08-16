import { inject, Injectable } from '@angular/core';
import type { Observable } from 'rxjs';
import { map } from 'rxjs';
import type { ProjectModel, ProjectCreateModel } from '../model/project.model';
import type { APIResponse } from '@shared/models';
import { HttpClient, HttpParams } from '@angular/common/http';

@Injectable()
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
