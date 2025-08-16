import {Component, inject} from '@angular/core';
import {filter, map, Observable, switchMap} from 'rxjs';
import {RecoveryFlow} from '@ory/kratos-client';
import {AsyncPipe} from '@angular/common';
import {FormMessagesComponent, KratosInputComponent, PreloaderComponent} from '@client/shared/ui';
import {ActivatedRoute, Params} from '@angular/router';
import {AuthService} from '../api';
import {environment} from '@client/environments/environment.development';
import {ApiResponse} from '@client/shared/models';

@Component({
    selector: "jr-recovery-form",
    imports: [
        FormMessagesComponent,
        KratosInputComponent,
        AsyncPipe,
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
export class RecoveryFormComponent {
    private route: ActivatedRoute = inject(ActivatedRoute);
    private authService: AuthService = inject(AuthService);
    flow$: Observable<RecoveryFlow | null> = this.route.queryParams.pipe(
        map((params: Params) => {
            if (!params.hasOwnProperty("flow")) {
                window.location.href = environment.AUTH_URL_RECOVERY_REDIRECT;
                return null;
            }
            return params["flow"];
        }),
        filter((flowId) => flowId !== null),
        switchMap((flowId: string) => this.authService.GetRecoveryForm(flowId).pipe(
            map((data: ApiResponse<RecoveryFlow>) => data.data)
        ))
    );
}
