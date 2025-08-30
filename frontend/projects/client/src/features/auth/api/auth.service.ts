import {inject, Injectable} from "@angular/core";
import {HttpClient, HttpErrorResponse, HttpParams} from "@angular/common/http";
import {catchError, EMPTY, finalize, Observable, shareReplay} from "rxjs";
import {FlowError, LoginFlow, RecoveryFlow, RegistrationFlow, VerificationFlow} from "@ory/kratos-client";
import {APIResponse} from '@client/shared/models';
import {UiService} from '@client/shared/ui';
import {environment, environment as env} from '@client/environments/environment.development'

export const AUTH_URLS = {
    LOGIN:          `${env.API_ROOT}/api/auth/form/login`,
    REGISTRATION:   `${env.API_ROOT}/api/auth/form/registration`,
    VERIFICATION:   `${env.API_ROOT}/api/auth/form/verification`,
    RECOVERY:       `${env.API_ROOT}/api/auth/form/recovery`,
    ERROR:          `${env.API_ROOT}/api/auth/form/error`,
}

@Injectable({ providedIn: "root" })
export class AuthService {
    http: HttpClient = inject(HttpClient)
    uiService: UiService = inject(UiService)

    public GetLoginForm(flow: string | null = null): Observable<APIResponse<LoginFlow>> {
        return this.getForm<LoginFlow>(AUTH_URLS.LOGIN, flow);
    }

    public GetRegistrationForm(flow: string | null = null): Observable<APIResponse<RegistrationFlow>> {
        return this.getForm<RegistrationFlow>(AUTH_URLS.REGISTRATION, flow);
    }

    public GetVerificationForm(flow: string | null = null): Observable<APIResponse<VerificationFlow>> {
        return this.getForm<VerificationFlow>(AUTH_URLS.VERIFICATION, flow);
    }

    public GetRecoveryForm(flow: string | null = null): Observable<APIResponse<RecoveryFlow>> {
        return this.getForm<RecoveryFlow>(AUTH_URLS.RECOVERY, flow);
    }

    public GetError(flow: string): Observable<APIResponse<FlowError>> {
        let params = new HttpParams().set("id", flow);
        this.uiService.loading.next(true);
        return this.http.get<APIResponse<FlowError>>(AUTH_URLS.ERROR, {params: params, withCredentials: true}).pipe(
            shareReplay(),
            finalize(() => this.uiService.loading.next(false)),
            catchError((error: HttpErrorResponse) => {
                console.error(error);
                return EMPTY;
            })
        );
    }

    private getForm<T>(url: string, flow: string | null = null): Observable<APIResponse<T>> {
        let params = new HttpParams();
        if (flow) params = params.append("flow", flow);
        this.uiService.loading.next(true);
        return this.http.get<APIResponse<T>>(url, {params: params, withCredentials: true}).pipe(
            shareReplay(),
            finalize(() => this.uiService.loading.next(false)),
            catchError((error: HttpErrorResponse) => {
                if (this.isStatusGone(error.status)) {
                    this.getNewFlowRedirect(url)
                    return EMPTY;
                }
                console.error(error);
                return EMPTY;
            })
        );
    }

    private isStatusGone(status: number): boolean {
        return status === 410;
    }

    private getNewFlowRedirect(url: string) {
        switch (url) {
            case AUTH_URLS.LOGIN:
                window.location.href = environment.AUTH_URL_LOGIN_REDIRECT;
                break;
            case AUTH_URLS.REGISTRATION:
                window.location.href = environment.AUTH_URL_REGISTRATION_REDIRECT;
                break;
            case AUTH_URLS.RECOVERY:
                window.location.href = environment.AUTH_URL_RECOVERY_REDIRECT;
                break;
        }
    }
}
