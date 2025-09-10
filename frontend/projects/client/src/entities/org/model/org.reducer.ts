import {OrgModel} from '@client/entities/org/model/org.model';
import {createReducer, on} from '@ngrx/store';
import {SetOrg} from '@client/entities/org/model/org.actions';

export interface OrgState {
    org: OrgModel | undefined
}

const initialState: OrgState = {
    org: undefined
}

export const orgReducer = createReducer(
    initialState,
    on(SetOrg, (state, action) => ({
        ...state,
        org: action.org,
    }))
)
