import { map, Observable } from 'rxjs';
import { UserModel } from '@entities/user';
import { inject, Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { APIResponse } from '@shared/models';

@Injectable({
    providedIn: 'root',
})
export class UserApi {
    http = inject(HttpClient);

    public GetUser(): Observable<UserModel> {
        return this.http
            .get<APIResponse<UserModel>>(`/api/v1/user/me`, { withCredentials: true })
            .pipe(map((response) => response.data));
    }
}
