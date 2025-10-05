import {Injectable} from "@angular/core";
import {map, Observable} from "rxjs";
import {environment as env} from '@client/environments/environment.development'
import {Project} from '@client/entities/project';
import {CommonApiService} from '@client/shared/common';

const projectUrl = {
    Root: `${env.API_ROOT}/api/v1/projects/`,
}

@Injectable({providedIn: "root"}) // TODO: does it has to be root?
export class ProjectService extends CommonApiService {
    public List(): Observable<Array<Project>> {
        return this.get<Array<Project>>(projectUrl.Root).pipe(
            map(data => data.data),
        );
    }

    public Create(payload: Record<string, string>): Observable<Project> {
        return this.post<Project>(projectUrl.Root, payload).pipe(
            map(data => data.data),
        );
    }
}
