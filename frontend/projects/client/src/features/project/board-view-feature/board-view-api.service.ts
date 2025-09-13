import {Injectable} from "@angular/core";
import {map, Observable} from "rxjs";
import {environment as env} from '@client/environments/environment.development'
import {CommonApiService} from '@client/shared/common';

const boardUrls = {
    TaskAction: (projectKey: string, taskID: string) => `${env.API_ROOT}/api/project/${projectKey}/tasks/${taskID}/status`,
}

interface ChangeStatusPayload {
    from_status: string
    to_status: string
    from_idx: number
    to_idx: number
}

@Injectable()
export class BoardViewApiService extends CommonApiService {
    public ChangeStatus(projectKey: string, taskCode: string, payload: ChangeStatusPayload): Observable<any> {
        return this.patch<any>(boardUrls.TaskAction(projectKey, taskCode), payload).pipe(
            map(data => data.data),
        );
    }
}
