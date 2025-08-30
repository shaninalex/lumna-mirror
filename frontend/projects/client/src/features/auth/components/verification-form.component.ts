import {Component, inject} from '@angular/core';
import {filter, map, Observable, switchMap} from 'rxjs';
import {VerificationFlow} from '@ory/kratos-client';
import {FormMessagesComponent, KratosInputComponent, PreloaderComponent} from '@client/shared/ui';
import {AsyncPipe} from '@angular/common';
import {ActivatedRoute, Params} from '@angular/router';
import {AuthService} from '../api';
import {APIResponse} from '@client/shared/models';


@Component({
    selector: "jr-verification-form",
    imports: [
        FormMessagesComponent,
        KratosInputComponent,
        AsyncPipe,
        PreloaderComponent
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
export class VerificationFormComponent {
    private route: ActivatedRoute = inject(ActivatedRoute);
    private authService: AuthService = inject(AuthService);
    flow$: Observable<VerificationFlow | null> = this.route.queryParams.pipe(
        map((params: Params) => {
            if (!params.hasOwnProperty("flow")) {
                return null;
            }
            return params["flow"];
        }),
        filter((flowID) => flowID !== null),
        switchMap((flowID: string) => this.authService.GetVerificationForm(flowID).pipe(
            map((data: APIResponse<VerificationFlow>) => data.data)
        ))
    );
}
