import {createFeatureSelector, createSelector} from '@ngrx/store';
import {OrgState} from '@client/entities/org/model/org.reducer';

const selectOrgFeature = createFeatureSelector<OrgState>('org');
export const selectOrg = createSelector(
    selectOrgFeature,
    (state: OrgState) => state.org
);
