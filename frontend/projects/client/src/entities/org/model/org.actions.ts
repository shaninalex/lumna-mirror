import {createAction, props} from '@ngrx/store';
import {OrgModel} from '@client/entities/org/model/org.model';

export const SetOrg = createAction(
    "[org] set",
    props<{ org: OrgModel }>(),
)
