import {catchError, EMPTY, finalize, Observable, shareReplay} from 'rxjs';
import {APIResponse} from '@client/shared/models';
import {HttpClient, HttpErrorResponse, HttpParams} from '@angular/common/http';
import {inject} from '@angular/core';
import {UiService} from '@client/shared/ui';


export class CommonApiService {
    http: HttpClient = inject(HttpClient)
    uiService: UiService = inject(UiService)

    get<T>(url: string, params?: HttpParams): Observable<APIResponse<T>> {
        this.uiService.loading.next(true);
        let p = new HttpParams()
        if (params) {
            p = params
        }
        return this.http.get<APIResponse<T>>(url, {params: p, withCredentials: true}).pipe(
            shareReplay(),
            finalize(() => this.uiService.loading.next(false)),
            catchError((error: HttpErrorResponse) => {
                console.error(error);
                return EMPTY;
            })
        );
    }

    patch<T>(url: string, data: any, params?: HttpParams): Observable<APIResponse<T>> {
        this.uiService.loading.next(true);
        let p = new HttpParams()
        if (params) {
            p = params
        }
        return this.http.patch<APIResponse<T>>(url, data, {params: p, withCredentials: true}).pipe(
            shareReplay(),
            finalize(() => this.uiService.loading.next(false)),
            catchError((error: HttpErrorResponse) => {
                console.error(error);
                return EMPTY;
            })
        );
    }
}
