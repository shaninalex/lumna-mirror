import {Component, inject, Input, OnInit} from '@angular/core';
import {FormsService} from '@client/entities/session';
import {filter, Observable, tap} from 'rxjs';
import {LoginFlow, RegistrationFlow} from '@ory/kratos-client';
import {KratosFormRenderer} from '@dev/ui/kratos';
import {LoaderComponent} from '@dev/ui/loader';
import {Store} from '@ngrx/store';
import {AppState} from '@client/shared/store';
import {RegisterFormLoadedAction} from '@client/features/auth';

@Component({
    selector: 'fr-registration-form',
    imports: [
        KratosFormRenderer,
        LoaderComponent
    ],
    template: `
        @if (loginForm$) {
            <ui-form-renderer [flow$]="loginForm$"/>
        } @else {
            <ui-loader/>
        }
    `
})
export class RegisterForm implements OnInit {
    @Input() flowID: string;
    service = inject(FormsService);
    loginForm$: Observable<RegistrationFlow>;
    private store: Store<AppState> = inject(Store<AppState>);

    ngOnInit() {
        this.loginForm$ = this.service.GetRegisterForm(this.flowID).pipe(
            filter(data => data !== undefined),
            tap(data => {
                this.store.dispatch(RegisterFormLoadedAction({form: data}))
            })
        )
    }
}
