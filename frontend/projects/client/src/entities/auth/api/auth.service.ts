import {inject, Injectable } from '@angular/core';
import {Observable} from 'rxjs';
import {HttpClient} from '@angular/common/http';
import {APIResponse} from '@client/shared/models';

@Injectable({
    providedIn: 'root'
})
export class AuthService {
    http = inject(HttpClient)

    login(data: any): Observable<APIResponse<any>> {
        return this.http.post<APIResponse<any>>("http://localhost:8000/api/auth/login", data, { withCredentials: true })
    }

    register(data: any): Observable<APIResponse<any>> {
        return this.http.post<APIResponse<any>>("http://localhost:8000/api/auth/register", data, { withCredentials: true })
    }
}
