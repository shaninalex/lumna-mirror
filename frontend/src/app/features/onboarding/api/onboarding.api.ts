import { inject, Injectable } from "@angular/core";
import { HttpClient } from "@angular/common/http";
import { TeamPageModel, WorkspacePageModel } from "../model";
import { map, Observable } from "rxjs";
import { APIResponse } from "@shared/models";
import { TeamOnboardingPage } from "@pages/onboarding/team";

@Injectable()
export class OnboardingApiService {
    http = inject(HttpClient);

    workspace(payload: WorkspacePageModel): Observable<any> {
        return this.http
            .post<
                APIResponse<any>
            >(`/api/v1/onboarding/workspace`, payload, { withCredentials: true })
            .pipe(map((response) => response.data));
    }

    team(payload: TeamPageModel): Observable<any> {
        return this.http
            .post<
                APIResponse<any>
            >(`/api/v1/onboarding/team`, payload, { withCredentials: true })
            .pipe(map((response) => response.data));
    }
}
