import { inject, Injectable } from "@angular/core";
import { HttpClient, HttpParams } from "@angular/common/http";
import { map, Observable } from "rxjs";
import { APIResponse } from "@shared/models";
import { OnboardingStateResponse, UserOnboardingModel } from "../model";
import { InvitationModel } from "@entities/invitation";

@Injectable()
export class OnboardingApiService {
    private http = inject(HttpClient);

    state(): Observable<OnboardingStateResponse> {
        return this.http
            .get<
                APIResponse<OnboardingStateResponse>
            >(`/api/v1/onboarding/state`, { withCredentials: true })
            .pipe(map((response) => response.data));
    }

    initialize(payload: UserOnboardingModel): Observable<InvitationModel> {
        return this.http
            .post<
                APIResponse<InvitationModel>
            >(`/api/v1/onboarding`, payload, { withCredentials: true })
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
