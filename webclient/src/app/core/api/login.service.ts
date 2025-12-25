import {inject, Injectable} from '@angular/core';
import {environment as env} from "@environments/environment.development"

import {LoginCredentials} from '@features/auth/login/model/login.model';
import {catchError, map, Observable, throwError} from 'rxjs';
import {HttpClient} from "@angular/common/http"
import {APIResponse} from '@shared/models';
import {UserModel} from '@entities/user';

export interface LoginApi {
    Login(credentials: LoginCredentials): Observable<any>;
}

@Injectable({
    providedIn: 'root',
})
export class LoginService {
    http = inject(HttpClient)

    public Login(payload: LoginCredentials): Observable<UserModel> {
        return this.http.post<APIResponse<UserModel>>(`${env.API_ROOT}/api/v1/auth/login`, payload, {withCredentials: true}).pipe(
            map(response => response.data)
        )
    }


    public Refresh(): Observable<void> {
        return this.http.get(`${env.API_ROOT}/api/v1/auth/refresh`, { withCredentials: true }).pipe(
                map(() => void 0),
                catchError((err) => throwError(() => err))
            );
    }
}
