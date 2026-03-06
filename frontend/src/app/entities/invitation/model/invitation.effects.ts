import { Injectable, inject } from '@angular/core';
import { Actions, createEffect, ofType } from '@ngrx/effects';
import { actionInvitationGetList, actionInvitationSetList } from './invitation.actions';
import { exhaustMap, map, of, switchMap } from 'rxjs';
import { InvitationApi } from '../api/invitation.api';

@Injectable()
export class InvitationEffects {
    private actions$ = inject(Actions);
    private invitationApi = inject(InvitationApi);

    get_list$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionInvitationGetList),
            exhaustMap(() =>
                this.invitationApi
                    .List()
                    .pipe(switchMap((invitations) => of(actionInvitationSetList({ invitations })))),
            ),
        ),
    );

    set_list$ = createEffect(() =>
        this.actions$.pipe(
            ofType(actionInvitationSetList),
            map((action) => actionInvitationSetList({ invitations: action.invitations })),
        ),
    );
}
