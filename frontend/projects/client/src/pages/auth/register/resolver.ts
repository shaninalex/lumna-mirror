import {inject} from '@angular/core';
import {ResolveFn} from '@angular/router';
import {catchError, of} from 'rxjs';
import {LoginFlow, RegistrationFlow} from '@ory/kratos-client';
import {FormsService} from '@client/entities/session';
import {environment} from '@client/environments/environment.development'


// login resolver
export const registrationFlowResolver: ResolveFn<RegistrationFlow | undefined> = (route) => {
    const flowID = route.queryParamMap.get('flow');
    if (!flowID) {
        window.location.href = environment.AUTH_URL_REGISTRATION_REDIRECT
        return of(undefined);
    }

    const service = inject(FormsService);
    return service.GetRegisterForm(flowID).pipe(
        catchError(() => of(undefined)) // handle errors gracefully
    );
};
