import {Injectable} from "@angular/core";
import {map, Observable} from "rxjs";
import {environment as env} from '@client/environments/environment.development'
import {CommonApiService} from '@client/shared/common';

const boardUrls = {
    changeTaskStatus: (taskID: number) => `${env.API_ROOT}/api/v1/task/${taskID}/status`,
}

export interface ChangeStatusPayload {
    from_status: number
    to_status: number
    from_idx: number
    to_idx: number
}

@Injectable({providedIn: 'root'})
export class BoardViewApiService extends CommonApiService {
    public ChangeStatus(taskId: number, payload: ChangeStatusPayload): Observable<any> {
        return this.patch<any>(boardUrls.changeTaskStatus(taskId), payload).pipe(
            map(data => data.data),
        );
    }
}
