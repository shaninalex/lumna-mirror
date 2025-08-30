import {Component, inject} from '@angular/core';
import {filter, map, Observable, of, switchMap} from 'rxjs';
import {RegistrationFlow} from '@ory/kratos-client';
import {AsyncPipe} from '@angular/common';
import {FormMessagesComponent, KratosInputComponent, PreloaderComponent} from '@client/shared/ui';
import {ActivatedRoute, Params} from '@angular/router';
import {AuthService} from '../api';
import {environment} from '@client/environments/environment.development';
import {APIResponse} from '@client/shared/models';


@Component({
    selector: "jr-registration-form",
    imports: [
        FormMessagesComponent,
        KratosInputComponent,
        PreloaderComponent,
        AsyncPipe
    ],
    template: `
        @if (flow$ | async; as flow) {
            <jr-ui-kratos-form-messages [messages]="flow.ui.messages"/>
            <form [action]="flow.ui.action" method="post" class="d-flex flex-column">
                @for (node of flow.ui.nodes; track node) {
                    <jr-ui-kratos-input [node]="node"/>
                }
            </form>
        } @else {
            <jr-preloader />
        }
    `
})
export class RegistrationFormComponent {
    private route: ActivatedRoute = inject(ActivatedRoute);
    private authService: AuthService = inject(AuthService);
    flow$: Observable<RegistrationFlow | null> = this.route.queryParams.pipe(
        map((params: Params) => {
            if (!params.hasOwnProperty("flow")) {
                window.location.href = environment.AUTH_URL_REGISTRATION_REDIRECT;
                return null;
            }
            return params["flow"];
        }),
        filter((flowID) => flowID !== null),
        switchMap((flowID: string) => this.authService.GetRegistrationForm(flowID).pipe(
            map((data: APIResponse<RegistrationFlow>) => data.data)
        ))
    );
}
