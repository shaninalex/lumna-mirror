import { inject, Injectable } from '@angular/core';
import { map, Observable } from 'rxjs';
import { UserModel } from '../model/user.model';
import { APIResponse } from '@shared/models';
import { HttpClient } from '@angular/common/http';

@Injectable({
    providedIn: 'root',
})
export class UserService {
    http = inject(HttpClient);

    getMe(): Observable<UserModel> {
        return this.http
            .get<APIResponse<UserModel>>(`/api/v1/user/me`, { withCredentials: true })
            .pipe(map((response) => response.data));
    }

    logout(): Observable<void> {
        return this.http
            .get<APIResponse<void>>(`/api/v1/user/logout`, { withCredentials: true })
            .pipe(map((response) => response.data));
    }
}
