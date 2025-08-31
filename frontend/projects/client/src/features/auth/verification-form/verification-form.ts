import {Component, inject, Input, OnInit} from '@angular/core';
import {FormsService} from '@client/entities/session';
import {filter, Observable, tap} from 'rxjs';
import {LoginFlow, VerificationFlow} from '@ory/kratos-client';
import {KratosFormRenderer} from '@dev/ui/kratos';
import {LoaderComponent} from '@dev/ui/loader';
import {Store} from '@ngrx/store';
import {AppState} from '@client/shared/store';
import {RegisterFormLoadedAction} from '@client/features/auth';

@Component({
    selector: 'fr-verification-form',
    imports: [
        KratosFormRenderer,
        LoaderComponent
    ],
    template: `
        @if (form$) {
            <ui-form-renderer [flow$]="form$"/>
        } @else {
            <ui-loader/>
        }
    `
})
export class VerificationForm implements OnInit {
    @Input() flowID: string;
    service = inject(FormsService);
    form$: Observable<VerificationFlow>;
    private store: Store<AppState> = inject(Store<AppState>);

    ngOnInit() {
        this.form$ = this.service.GetVerificationForm(this.flowID).pipe(
            filter(data => data !== undefined),
            tap(data => {
                this.store.dispatch(RegisterFormLoadedAction({form: data}))
            })
        )
    }
}
