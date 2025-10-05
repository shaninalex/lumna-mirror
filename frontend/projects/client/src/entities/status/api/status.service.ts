import {Injectable} from "@angular/core";
import {map, Observable} from "rxjs";
import {environment as env} from '@client/environments/environment.development'
import {CommonApiService} from '@client/shared/common';
import {Status, StatusInput} from '@client/entities/status/model';

const statusUrl = {
    statusRoot: (projectId: number) => `${env.API_ROOT}/api/v1/project/${projectId}/statuses`,
    statusSort: (projectId: number) => `${env.API_ROOT}/api/v1/project/${projectId}/statuses-sort`,
    statusItem:  (projectId: number, statusId: number) => `${env.API_ROOT}/api/v1/project/${projectId}/statuses/${statusId}`,
}

@Injectable({providedIn: "root"})
export class StatusService extends CommonApiService {
    public List(projectId: number): Observable<Status[]> {
        return this.get<Status[]>(statusUrl.statusRoot(projectId)).pipe(
            map(data => data.data),
        );
    }

    public Create(projectId: number, payload: StatusInput): Observable<Status> {
        return this.post<Status>(statusUrl.statusRoot(projectId), payload).pipe(
            map(data => data.data),
        );
    }

    public Patch(projectId: number, statusId: number, payload: StatusInput): Observable<Status> {
        return this.patch<Status>(statusUrl.statusItem(projectId, statusId), payload).pipe(
            map(data => data.data),
        );
    }

    public Delete(projectId: number, statusId: number): Observable<null> {
        return this.delete<null>(statusUrl.statusItem(projectId, statusId)).pipe(
            map(data => data.data),
        );
    }

    public StatusSort(projectId: number, payload: Record<number, number>): Observable<Status[]> {
        return this.patch<Status[]>(statusUrl.statusSort(projectId), payload).pipe(
            map(data => data.data),
        );
    }
}
