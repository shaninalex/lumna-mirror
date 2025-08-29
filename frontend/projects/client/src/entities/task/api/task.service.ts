import {Injectable} from "@angular/core";
import {HttpParams} from "@angular/common/http";
import {map, Observable} from "rxjs";
import {environment as env} from '@client/environments/environment.development'
import {Task} from '@client/entities/task';
import {CommonApiService} from '@client/shared/common';

export const TASKS_URLS = {
    Detail: (projectCode: string, taskCode: string) => `${env.API_ROOT}/api/project/${projectCode}/tasks/${taskCode}`
}

@Injectable({providedIn: "root"})
export class TaskService extends CommonApiService {
    public Detail(projectCode: string, taskCode: string): Observable<Task> {
        let p = new HttpParams()
        return this.get<Task>(TASKS_URLS.Detail(projectCode, taskCode), p).pipe(
            map(resp => resp.data)
        );
    }

    public Update(projectCode: string, taskCode: string, data: any): Observable<Task> {
        return this.patch<Task>(TASKS_URLS.Detail(projectCode, taskCode), data).pipe(
            map(resp => resp.data)
        );
    }
}
