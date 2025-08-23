import {inject, Injectable} from "@angular/core";
import {HttpClient, HttpErrorResponse} from "@angular/common/http";
import {catchError, EMPTY, finalize, map, Observable, shareReplay} from "rxjs";
import {ApiResponse} from '@client/shared/models';
import {UiService} from '@client/shared/ui';
import {environment as env} from '@client/environments/environment.development'
import {Project} from '@client/entities/project';

export const PROJECT_URLS = {
    List: `${env.API_ROOT}/api/project/list`,
}

@Injectable({providedIn: "root"})
export class ProjectService {
    http: HttpClient = inject(HttpClient)
    uiService: UiService = inject(UiService)

    public List(): Observable<Array<Project>> {
        return this.getForm<Array<Project>>(PROJECT_URLS.List).pipe(
            map(data => data.data),
        );
    }

    private getForm<T>(url: string): Observable<ApiResponse<T>> {
        this.uiService.loading.next(true);
        return this.http.get<ApiResponse<T>>(url, {withCredentials: true}).pipe(
            shareReplay(),
            finalize(() => this.uiService.loading.next(false)),
            catchError((error: HttpErrorResponse) => {
                console.error(error);
                return EMPTY;
            })
        );
    }
}
