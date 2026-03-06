import { createAction, props } from '@ngrx/store';
import { InvitationModel } from './invitation.model';

export const actionInvitationGetList = createAction(
    '[Invitation] get list',
    props<{ boardId: number }>(),
);

export const actionInvitationSetList = createAction(
    '[Invitation] set list',
    props<{ invitations: InvitationModel[] }>(),
);
