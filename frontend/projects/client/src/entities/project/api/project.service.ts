import {Injectable} from "@angular/core";
import {map, Observable} from "rxjs";
import {environment as env} from '@client/environments/environment.development'
import {Project} from '@client/entities/project';
import {CommonApiService} from '@client/shared/common';

export const PROJECT_URLS = {
    List: `${env.API_ROOT}/api/project/list`,
}

@Injectable({providedIn: "root"}) // TODO: does it has to be root?
export class ProjectService extends CommonApiService {
    public List(): Observable<Array<Project>> {
        return this.get<Array<Project>>(PROJECT_URLS.List).pipe(
            map(data => data.data),
        );
    }
}
