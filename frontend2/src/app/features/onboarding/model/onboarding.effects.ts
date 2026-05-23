import { Actions, createEffect, ofType } from "@ngrx/effects";
import { inject, Injectable } from "@angular/core";
import {
    actionOnboardingCreateInvite,
    actionOnboardingCreateInviteFailed,
    actionOnboardingCreateInviteSuccess,
    actionOnboardingValidateInviteFailed,
    actionOnboardingValidateInviteSuccess,
    actionOnboardingValidateInviteToken
} from "./onboarding.actions";
import { catchError, exhaustMap, map, of, tap } from "rxjs";
import { OnboardingApiService } from "../api/onboarding.api";
import { Error } from "@shared/models";

@Injectable()
export class OnboardingEffects {
    private actions$ = inject(Actions);
    private api = inject(OnboardingApiService);

    invitation_init$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionOnboardingValidateInviteToken),
            exhaustMap((action) =>
                this.api.invitationValidateToken(action.token).pipe(
                    map((invitation) =>
                        actionOnboardingValidateInviteSuccess({ invitation })
                    ),
                    catchError((err) =>
                        of(
                            actionOnboardingValidateInviteFailed({
                                error: err as Error
                            })
                        )
                    )
                )
            )
        )
    );

    onboarding_init$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionOnboardingCreateInvite),
            exhaustMap((action) =>
                this.api.initialize(action.payload).pipe(
                    map((invitation) =>
                        actionOnboardingCreateInviteSuccess({ invitation })
                    ),
                    catchError((err) =>
                        of(
                            actionOnboardingCreateInviteFailed({
                                error: err as Error[]
                            })
                        )
                    )
                )
            )
        )
    );
}
