import {inject, Injectable} from '@angular/core';
import {environment as env} from "@environments/environment.development"

import {LoginCredentials} from '@features/auth/login/model/login.model';
import {map, Observable} from 'rxjs';
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
}
