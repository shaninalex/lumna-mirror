import {inject, Injectable} from '@angular/core';
import {map, Observable} from 'rxjs';
import {environment} from '@client/environments/environment.development';
import {HttpClient} from '@angular/common/http';
import {Settings, UserModel} from '@client/entities/user';
import {APIResponse} from '@client/shared/models';

@Injectable({ providedIn: 'root' })
export class UserService {
    http = inject(HttpClient);

    getUser(): Observable<UserModel> {
        return this.http.get<APIResponse<UserModel>>(`${environment.API_ROOT}/api/user/me`, { withCredentials: true }).pipe(
            map(data => data.data),
        );
    }

    updateUserSettings(settings: Settings): Observable<UserModel> {
        return this.http.post<APIResponse<UserModel>>(`${environment.API_ROOT}/api/user/settings`, settings, { withCredentials: true }).pipe(
            map(data => data.data),
        );
    }

    logout(): Observable<APIResponse<any>> {
        return this.http.get<APIResponse<any>>(`${environment.API_ROOT}/api/user/logout`, { withCredentials: true })
    }
}
