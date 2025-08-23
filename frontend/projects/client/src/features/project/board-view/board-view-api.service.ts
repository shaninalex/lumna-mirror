import {Injectable} from "@angular/core";
import {map, Observable} from "rxjs";
import {environment as env} from '@client/environments/environment.development'
import {Status} from '@client/entities/project';
import {CommonApiService} from '@client/shared/common';

export const BOARD_URLS = {
    Statuses: (projectKey: string) =>  `${env.API_ROOT}/api/project/${projectKey}/statuses`,
}

@Injectable()
export class BoardViewApiService extends CommonApiService {
    public Statuses(projectKey: string): Observable<Status[]> {
        return this.getForm<Status[]>(BOARD_URLS.Statuses(projectKey)).pipe(
            map(data => data.data),
        );
    }
}
