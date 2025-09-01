import {Injectable} from "@angular/core";
import {map, Observable} from "rxjs";
import {environment as env} from '@client/environments/environment.development'
import {CommonApiService} from '@client/shared/common';
import {Task} from '@client/entities/task';

const tasksUrl = {
    list: (projectCode: string) => `${env.API_ROOT}/api/project/${projectCode}/tasks`,
}

@Injectable({providedIn: "root"}) // TODO: does it has to be root?
export class TaskService extends CommonApiService {
    public List(projectCode: string): Observable<Task[]> {
        return this.get<Task[]>(tasksUrl.list(projectCode)).pipe(
            map(data => data.data),
        );
    }
}
