import {inject} from '@angular/core';
import {ResolveFn} from '@angular/router';
import {catchError, of} from 'rxjs';
import {LoginFlow} from '@ory/kratos-client';
import {FormsService} from '@client/entities/session';
import {environment} from '@client/environments/environment.development'


// login resolver
export const loginFlowResolver: ResolveFn<LoginFlow | undefined> = (route) => {
    const flowID = route.queryParamMap.get('flow');
    if (!flowID) {
        window.location.href = environment.AUTH_URL_LOGIN_REDIRECT
        return of(undefined);
    }

    const service = inject(FormsService);
    return service.GetLoginForm(flowID).pipe(
        catchError(() => of(undefined)) // handle errors gracefully
    );
};
