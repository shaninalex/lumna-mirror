import { inject, Injectable } from "@angular/core";
import { HttpClient, HttpParams } from "@angular/common/http";
import { map, Observable } from "rxjs";
import { APIResponse } from "@shared/models";
import { OnboardingState, UserOnboardingModel } from "../model";

@Injectable()
export class OnboardingApiService {
    private http = inject(HttpClient);

    state(): Observable<OnboardingState> {
        return this.http
            .get<
                APIResponse<any>
            >(`/api/v1/onboarding/state`, { withCredentials: true })
            .pipe(map((response) => response.data));
    }

    user(payload: UserOnboardingModel): Observable<any> {
        return this.http
            .post<
                APIResponse<any>
            >(`/api/v1/onboarding/user`, payload, { withCredentials: true })
            .pipe(map((response) => response.data));
    }

    invitationValidateToken(token: string): Observable<any> {
        const params = new HttpParams().append("token", token);
        return this.http
            .get<
                APIResponse<any>
            >(`/api/v1/onboarding/invitation`, { params, withCredentials: true })
            .pipe(map((response) => response.data));
    }
}
