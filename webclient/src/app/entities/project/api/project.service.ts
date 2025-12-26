import {inject, Injectable} from '@angular/core';
import {map, Observable} from 'rxjs';
import {ProjectModel, ProjectPayload} from '@entities/project';
import {APIResponse} from '@shared/models';
import {environment as env} from '@environments/environment.development';
import {HttpClient} from '@angular/common/http';

@Injectable({
    providedIn: 'root',
})
export class ProjectService {
    http = inject(HttpClient)

    GetProjects(): Observable<ProjectModel[]> {
        return this.http.get<APIResponse<ProjectModel[]>>(`${env.API_ROOT}/api/v1/projects`, {withCredentials: true}).pipe(
            map(response => response.data)
        )
    }

    CreateProject(payload: ProjectPayload): Observable<ProjectModel> {
        return this.http.post<APIResponse<ProjectModel>>(`${env.API_ROOT}/api/v1/projects`, payload, {withCredentials: true}).pipe(
            map(response => response.data)
        )
    }

    DeleteProject(projectId: number): Observable<void> {
        return this.http.delete<APIResponse<void>>(`${env.API_ROOT}/api/v1/project/${projectId}`, {withCredentials: true}).pipe(
            map(response => response.data)
        )
    }
}
