import { createEntityAdapter, EntityState } from '@ngrx/entity';
import { createReducer, on } from '@ngrx/store';
import { actionInvitationSetList } from './invitation.actions';
import { InvitationModel } from './invitation.model';

export interface InvitationState extends EntityState<InvitationModel> {}
export const invitationAdapter = createEntityAdapter<InvitationModel>();
const initialState = invitationAdapter.getInitialState();

export const invitationReducer = createReducer(
    initialState,

    on(actionInvitationSetList, (state, { invitations }) =>
        invitationAdapter.addMany(invitations, state),
    ),
);
