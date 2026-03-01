import {createAction, props} from '@ngrx/store';
import {ActivityModel} from './activity.model';

export const actionActivityGetList = createAction(
    '[Activity] get list',
    props<{ entity_type: string, entity_id: number }>()
);

export const actionActivitySetList = createAction(
    '[Activity] set list',
    props<{ list: ActivityModel[] }>()
);
