import {inject, Injectable} from "@angular/core";
import {HttpClient, HttpErrorResponse, HttpParams} from "@angular/common/http";
import {catchError, EMPTY, finalize, Observable, shareReplay} from "rxjs";
import {ApiResponse} from '@client/shared/models';
import {UiService} from '@client/shared/ui';
import {environment as env} from '@client/environments/environment.development'

export const TASKS_URLS = {
    List: `${env.API_ROOT}/api/project/tasks`,
}

@Injectable({providedIn: "root"})
export class TaskService {
    http: HttpClient = inject(HttpClient)
    uiService: UiService = inject(UiService)

    public List(projectKey: string): Observable<ApiResponse<any>> {
        let p = new HttpParams()
        p = p.append("project", projectKey)
        return this.getForm<any>(TASKS_URLS.List, p);
    }

    private getForm<T>(url: string, params: HttpParams): Observable<ApiResponse<T>> {
        this.uiService.loading.next(true);
        return this.http.get<ApiResponse<T>>(url, {params: params, withCredentials: true}).pipe(
            shareReplay(),
            finalize(() => this.uiService.loading.next(false)),
            catchError((error: HttpErrorResponse) => {
                console.error(error);
                return EMPTY;
            })
        );
    }
}
