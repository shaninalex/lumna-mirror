import {Injectable} from "@angular/core";
import {map, Observable} from "rxjs";
import {environment as env} from '@client/environments/environment.development'
import {Status} from '@client/entities/project';
import {CommonApiService} from '@client/shared/common';

export const BOARD_URLS = {
    Statuses: (projectKey: string) => `${env.API_ROOT}/api/project/${projectKey}/statuses`,
    TaskAction: (projectKey: string, taskID: string) => `${env.API_ROOT}/api/project/${projectKey}/task/${taskID}/status`,
}

interface ChangeStatusPayload {
    from_status: string
    to_status: string
    from_idx: number
    to_idx: number
}

@Injectable()
export class BoardViewApiService extends CommonApiService {
    public Statuses(projectKey: string): Observable<Status[]> {
        return this.get<Status[]>(BOARD_URLS.Statuses(projectKey)).pipe(
            map(data => data.data),
        );
    }

    public ChangeStatus(projectKey: string, taskID: string, payload: ChangeStatusPayload): Observable<any> {
        return this.patch<any>(BOARD_URLS.TaskAction(projectKey, taskID), payload).pipe(
            map(data => data.data),
        );
    }
}
