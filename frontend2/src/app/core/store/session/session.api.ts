import { inject, Injectable } from "@angular/core";
import { HttpClient } from "@angular/common/http";
import { finalize, map, Observable, shareReplay } from "rxjs";
import { UserModel } from "@entities/user";
import { APIResponse } from "@shared/models";

@Injectable()
export class SessionApi {
    private http = inject(HttpClient);
    private refreshRequest$: Observable<void> | null = null;

    login(email: string, password: string): Observable<UserModel> {
        return this.http
            .post<
                APIResponse<UserModel>
            >(`/api/v1/auth/login`, { email, password }, { withCredentials: true })
            .pipe(map((response) => response.data));
    }

    logout(): Observable<void> {
        return this.http
            .get<
                APIResponse<void>
            >(`/api/v1/auth/logout`, { withCredentials: true })
            .pipe(map((response) => response.data));
    }

    refresh(): Observable<void> {
        if (!this.refreshRequest$) {
            this.refreshRequest$ = this.http
                .get<void>("/api/v1/auth/refresh", {
                    withCredentials: true
                })
                .pipe(
                    shareReplay(1),
                    finalize(() => {
                        this.refreshRequest$ = null;
                    })
                );
        }

        return this.refreshRequest$;
    }
}
