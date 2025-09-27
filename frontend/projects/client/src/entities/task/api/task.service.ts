import {Injectable} from "@angular/core";
import {map, Observable} from "rxjs";
import {environment as env} from '@client/environments/environment.development'
import {CommonApiService} from '@client/shared/common';
import {CreateTaskDto, Task} from '@client/entities/task';

const tasksUrl = {
    tasksRoot: (projectCode: string) => `${env.API_ROOT}/api/v1/project/${projectCode}/tasks`,
    tasksCreate: `${env.API_ROOT}/api/v1/project/tasks`,
}

@Injectable({providedIn: "root"}) // TODO: does it has to be root?
export class TaskService extends CommonApiService {
    public List(projectCode: string): Observable<Task[]> {
        return this.get<Task[]>(tasksUrl.tasksRoot(projectCode)).pipe(
            map(data => data.data),
        );
    }

    public Create(payload: CreateTaskDto): Observable<Task> {
        return this.post<Task>(tasksUrl.tasksCreate, payload).pipe(
            map(data => data.data),
        );
    }
}
