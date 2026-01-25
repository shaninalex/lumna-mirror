import { inject, Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { map, Observable } from 'rxjs';
import { UserModel } from '@entities/user';
import { APIResponse } from '@shared/models';

@Injectable({
    providedIn: 'root',
})
export class AuthService {
    http = inject(HttpClient);

    login(email: string, password: string): Observable<UserModel> {
        return this.http
            .post<APIResponse<UserModel>>(
                `/api/v1/auth/login`,
                { email, password },
                { withCredentials: true }
            )
            .pipe(map((response) => response.data));
    }
}
