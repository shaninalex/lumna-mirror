import {Component, inject} from '@angular/core';
import {filter, map, Observable, switchMap} from 'rxjs';
import {LoginFlow} from '@ory/kratos-client';
import {FormMessagesComponent, KratosInputComponent, PreloaderComponent} from '@client/shared/ui';
import {AsyncPipe} from '@angular/common';
import {ActivatedRoute, Params} from '@angular/router';
import {AuthService} from '../api';
import {ApiResponse} from '@client/shared/models';
import {environment} from '@client/environments/environment.development';

@Component({
    selector: "jr-login-form",
    imports: [
        FormMessagesComponent,
        KratosInputComponent,
        AsyncPipe,
        PreloaderComponent
    ],
    template: `
        @if (loginFlow$ | async; as loginFlow) {
            <jr-ui-kratos-form-messages [messages]="loginFlow.ui.messages"/>
            <form [action]="loginFlow.ui.action" method="post" class="d-flex flex-column">
                @for (node of loginFlow.ui.nodes; track node) {
                    <jr-ui-kratos-input [node]="node"/>
                }
            </form>
        } @else {
            <jr-preloader />
        }
    `
})
export class LoginFormComponent {
    private route: ActivatedRoute = inject(ActivatedRoute);
    private authService: AuthService = inject(AuthService);
    loginFlow$: Observable<LoginFlow | null> = this.route.queryParams.pipe(
        map((params: Params) => {
            if (!params.hasOwnProperty("flow")) {
                window.location.href = environment.AUTH_URL_LOGIN_REDIRECT;
                return null;
            }
            return params["flow"];
        }),
        filter((flowId) => flowId !== null),
        switchMap((flowId: string) => this.authService.GetLoginForm(flowId).pipe(
            map((data: ApiResponse<LoginFlow>) => data.data)
        ))
    );
}
