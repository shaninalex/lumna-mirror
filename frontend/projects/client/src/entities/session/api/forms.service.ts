import {Injectable} from '@angular/core';
import {HttpErrorResponse, HttpParams} from '@angular/common/http';
import {catchError, EMPTY, map, Observable} from 'rxjs';
import {LoginFlow, RegistrationFlow, VerificationFlow} from '@ory/kratos-client';
import {CommonApiService} from '@client/shared/common';
import {environment} from '@client/environments/environment.development';

const formUrls = {
    login: `${environment.API_ROOT}/api/auth/form/login`,
    register: `${environment.API_ROOT}/api/auth/form/registration`,
    verification: `${environment.API_ROOT}/api/auth/form/verification`,
}

const OnFlowGone = (error: HttpErrorResponse) => {
    if (error.status === 410) {
        window.location.href = environment.AUTH_URL_LOGIN_REDIRECT;
        return EMPTY;
    }
    return EMPTY;
}

@Injectable({
    providedIn: 'root'
})
export class FormsService extends CommonApiService {
    public GetLoginForm(flow: string): Observable<LoginFlow> {
        let p = new HttpParams()
        p = p.append("flow", flow)
        return this.get<LoginFlow>(formUrls.login, p).pipe(
            map(resp => resp.data),
            catchError(OnFlowGone)
        );
    }

    public GetRegisterForm(flow: string): Observable<RegistrationFlow> {
        let p = new HttpParams()
        p = p.append("flow", flow)
        return this.get<RegistrationFlow>(formUrls.register, p).pipe(
            map(resp => resp.data),
            catchError(OnFlowGone)
        );
    }

    public GetVerificationForm(flow: string): Observable<VerificationFlow> {
        let p = new HttpParams()
        p = p.append("flow", flow)
        return this.get<VerificationFlow>(formUrls.verification, p).pipe(
            map(resp => resp.data),
            catchError(OnFlowGone)
        );
    }
}
