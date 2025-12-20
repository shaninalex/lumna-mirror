import {inject, Injectable} from '@angular/core';
import { environment as env } from "@environments/environment.development"

import {LoginCredentials} from '@features/auth/login/model/login.model';
import {Observable, of, timeout} from 'rxjs';
import { HttpClient } from "@angular/common/http"
@Injectable({
    providedIn: 'root',
})
export class LoginService {
    http = inject(HttpClient)

    public Login(payload: LoginCredentials): Observable<any> {
        return this.http.post<any>(`${env.API_ROOT}/api/v1/auth/login`, payload, { withCredentials: true })
    }
}
