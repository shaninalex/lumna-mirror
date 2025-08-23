import {catchError, EMPTY, finalize, Observable, shareReplay} from 'rxjs';
import {ApiResponse} from '@client/shared/models';
import {HttpClient, HttpErrorResponse} from '@angular/common/http';
import {inject} from '@angular/core';
import {UiService} from '@client/shared/ui';


export class CommonApiService {
    http: HttpClient = inject(HttpClient)
    uiService: UiService = inject(UiService)

    getForm<T>(url: string): Observable<ApiResponse<T>> {
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
