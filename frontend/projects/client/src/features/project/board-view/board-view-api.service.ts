import {Injectable} from "@angular/core";
import {map, Observable} from "rxjs";
import {environment as env} from '@client/environments/environment.development'
import {Status} from '@client/entities/project';
import {CommonApiService} from '@client/shared/common';

export const BOARD_URLS = {
    Tasks: (projectKey: string) => `${env.API_ROOT}/api/project/${projectKey}/tasks`,
    TaskAction: (projectCode: string, taskCode: string) => `${env.API_ROOT}/api/project/${projectCode}/tasks/${taskCode}/status`,
}

interface ChangeStatusPayload {
    from_status: string
    to_status: string
    from_idx: number
    to_idx: number
}

@Injectable()
export class BoardViewApiService extends CommonApiService {
    public Tasks(projectKey: string): Observable<Status[]> {
        return this.get<Status[]>(BOARD_URLS.Tasks(projectKey)).pipe(
            map(data => data.data),
        );
    }

    public ChangeStatus(projectCode: string, taskCode: string, payload: ChangeStatusPayload): Observable<any> {
        return this.patch<any>(BOARD_URLS.TaskAction(projectCode, taskCode), payload).pipe(
            map(data => data.data),
        );
    }
}
