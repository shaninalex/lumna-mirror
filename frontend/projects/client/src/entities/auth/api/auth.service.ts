import {inject, Injectable} from '@angular/core';
import {environment as env} from '@client/environments/environment.development';
import {Observable} from 'rxjs';
import {HttpClient} from '@angular/common/http';
import {APIResponse} from '@client/shared/models';

@Injectable({providedIn: 'root'})
export class AuthService {
    http = inject(HttpClient)

    login(data: any): Observable<APIResponse<any>> {
        return this.http.post<APIResponse<any>>(`${env.API_ROOT}/api/auth/login`, data, {withCredentials: true})
    }

    register(data: any): Observable<APIResponse<any>> {
        return this.http.post<APIResponse<any>>(`${env.API_ROOT}/api/auth/register`, data, {withCredentials: true})
    }
}
