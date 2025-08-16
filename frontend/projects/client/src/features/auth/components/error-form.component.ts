import {Component, inject} from '@angular/core';
import {filter, map, Observable, switchMap} from 'rxjs';
import {FlowError} from '@ory/kratos-client';
import {ApiResponse} from '@client/shared/models';
import {AsyncPipe, JsonPipe} from '@angular/common';
import {ActivatedRoute, Params} from '@angular/router';
import {AuthService} from '../api';
import {PreloaderComponent} from '@client/shared/ui';

@Component({
    selector: "jr-error-form",
    imports: [
        AsyncPipe,
        JsonPipe,
        PreloaderComponent
    ],
    template: `
        @if (flow$ | async; as flow) {
            {{ flow.error | json }}
        } @else {
            <jr-preloader />
        }
    `
})
export class ErrorFormComponent {
    private route: ActivatedRoute = inject(ActivatedRoute);
    private authService: AuthService = inject(AuthService);
    flow$: Observable<FlowError | null> = this.route.queryParams.pipe(
        map((params: Params) => {
            if (!params.hasOwnProperty("flow")) {
                return null;
            }
            return params["flow"];
        }),
        filter((flowId) => flowId !== null),
        switchMap((flowId: string) => this.authService.GetError(flowId).pipe(
            map((data: ApiResponse<FlowError>) => data.data)
        ))
    );
}
