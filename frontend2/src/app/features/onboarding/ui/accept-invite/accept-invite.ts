import { Component, inject, Input, OnInit } from "@angular/core";
import { InvitationProcessState } from "@features/onboarding";
import {
    actionOnboardingValidateInviteFailed,
    actionOnboardingValidateInviteSuccess,
    actionOnboardingValidateInviteToken
} from "@features/onboarding/model/onboarding.actions";

import { Actions, ofType } from "@ngrx/effects";
import { Store } from "@ngrx/store";
import { tap } from "rxjs";

@Component({
    selector: "app-accept-invite-feature",
    template: ``
})
export class AcceptInviteFeature implements OnInit {
    @Input() token: string;
    process: InvitationProcessState = InvitationProcessState.VALIDATING;

    private store = inject(Store);
    private actions$ = inject(Actions);

    ngOnInit(): void {
        this.store.dispatch(
            actionOnboardingValidateInviteToken({ token: this.token })
        );

        this.actions$
            .pipe(
                ofType(
                    actionOnboardingValidateInviteSuccess,
                    actionOnboardingValidateInviteFailed
                ),
                tap((action) => {
                    console.log(action);
                })
            )
            .subscribe();
    }
}
