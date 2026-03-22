import { createFeatureSelector, createSelector } from '@ngrx/store';
import { invitationAdapter, InvitationState } from './invitation.store';

const selectInvitationFeature = createFeatureSelector<InvitationState>('column');

const invitationSelectors = invitationAdapter.getSelectors();

export const selectInvitations = createSelector(selectInvitationFeature, (state) =>
    invitationSelectors.selectAll(state),
);
