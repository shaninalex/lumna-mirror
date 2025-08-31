import {inject} from '@angular/core';
import {ResolveFn} from '@angular/router';
import {catchError, of} from 'rxjs';
import {VerificationFlow} from '@ory/kratos-client';
import {FormsService} from '@client/entities/session';


// login resolver
export const verificationFlowResolver: ResolveFn<VerificationFlow | undefined> = (route) => {
    const flowID = route.queryParamMap.get('flow');
    if (!flowID) {
        return of(undefined);
    }

    const service = inject(FormsService);
    return service.GetVerificationForm(flowID).pipe(
        catchError(() => of(undefined)) // handle errors gracefully
    );
};
