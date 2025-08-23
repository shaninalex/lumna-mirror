import {Injectable} from "@angular/core";
import {map, Observable} from "rxjs";
import {environment as env} from '@client/environments/environment.development'
import {Project} from '@client/entities/project';
import {CommonApiService} from '@client/shared/common';

export const BOARD_URLS = {
    BOARD: (projectKey: string) =>  `${env.API_ROOT}/api/project/${projectKey}/board`,
}

@Injectable()
export class BoardViewApiService extends CommonApiService {
    public BoardView(projectKey: string): Observable<Array<Project>> {
        return this.getForm<Array<Project>>(BOARD_URLS.BOARD(projectKey)).pipe(
            map(data => data.data),
        );
    }
}
