import {Observable, shareReplay} from 'rxjs';
import {APIResponse} from '@client/shared/models';
import {HttpClient, HttpParams} from '@angular/common/http';
import {inject} from '@angular/core';


export class CommonApiService {
    http: HttpClient = inject(HttpClient)

    get<T>(url: string, params?: HttpParams): Observable<APIResponse<T>> {
        let p = new HttpParams()
        if (params) {
            p = params
        }
        return this.http.get<APIResponse<T>>(url, {params: p, withCredentials: true}).pipe(
            shareReplay(),
        );
    }

    patch<T>(url: string, data: any, params?: HttpParams): Observable<APIResponse<T>> {
        let p = new HttpParams()
        if (params) {
            p = params
        }
        return this.http.patch<APIResponse<T>>(url, data, {params: p, withCredentials: true}).pipe(
            shareReplay(),
        );
    }

    post<T>(url: string, data: any, params?: HttpParams): Observable<APIResponse<T>> {
        let p = new HttpParams()
        if (params) {
            p = params
        }
        return this.http.post<APIResponse<T>>(url, data, {params: p, withCredentials: true}).pipe(
            shareReplay(),
        );
    }

    delete<T>(url: string): Observable<APIResponse<T>> {
        return this.http.delete<APIResponse<T>>(url, {withCredentials: true}).pipe(
            shareReplay(),
        );
    }
}
