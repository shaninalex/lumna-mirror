import {Injectable} from "@angular/core";
import {HttpParams} from "@angular/common/http";
import {map, Observable} from "rxjs";
import {environment as env} from '@client/environments/environment.development'
import {Task} from '@client/entities/task';
import {CommonApiService} from '@client/shared/common';

export const TASKS_URLS = {
    List: `${env.API_ROOT}/api/project/tasks`,
}

@Injectable({providedIn: "root"})
export class TaskService extends CommonApiService {
    public List(projectKey: string): Observable<Task[]> {
        let p = new HttpParams()
        p = p.append("project", projectKey)
        return this.get<Task[]>(TASKS_URLS.List, p).pipe(
            map(resp => resp.data)
        );
    }
}
