import {
    Component,
    DestroyRef,
    inject,
    Input,
    OnInit,
    signal
} from "@angular/core";
import { takeUntilDestroyed } from "@angular/core/rxjs-interop";
import { InvitationProcessState } from "@features/onboarding";
import {
    actionOnboardingValidateInviteFailed,
    actionOnboardingValidateInviteSuccess,
    actionOnboardingValidateInviteToken
} from "@features/onboarding/model/onboarding.actions";

import { Actions, ofType } from "@ngrx/effects";
import { Store } from "@ngrx/store";
import { tap } from "rxjs";
import { RegisterForm } from "@features/auth";

@Component({
    selector: "app-accept-invite-feature",
    template: `
        @switch (process()) {
            @case (InvitationProcessState.PENDING) {
                <div>Loading...</div>
            }
            @case (InvitationProcessState.FAILED) {
                <div class="text-red-500">
                    Invlid or expired invitation link
                </div>
            }
            @case (InvitationProcessState.SUCCESS) {
                <app-register-form-feature [invitationToken]="token" />
            }
        }
    `,
    imports: [RegisterForm]
})
export class AcceptInviteFeature implements OnInit {
    @Input() token: string;
    process = signal<InvitationProcessState>(InvitationProcessState.PENDING);

    public readonly InvitationProcessState: typeof InvitationProcessState =
        InvitationProcessState;
    private destroyRef = inject(DestroyRef);

    private store = inject(Store);
    private actions$ = inject(Actions);

    ngOnInit(): void {
        this.store.dispatch(
            actionOnboardingValidateInviteToken({ token: this.token })
        );

        this.actions$
            .pipe(
                ofType(actionOnboardingValidateInviteFailed),
                takeUntilDestroyed(this.destroyRef),
                tap(() => this.process.set(InvitationProcessState.FAILED))
            )
            .subscribe();

        this.actions$
            .pipe(
                ofType(actionOnboardingValidateInviteSuccess),
                takeUntilDestroyed(this.destroyRef),
                tap(() => this.process.set(InvitationProcessState.SUCCESS))
            )
            .subscribe();
    }
}
