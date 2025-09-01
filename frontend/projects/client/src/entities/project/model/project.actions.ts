import {createAction, props} from '@ngrx/store';
import {Project} from '@client/entities/project/model/project.model';

export const SetProjectsAction = createAction(
    "[project] set list",
    props<{ payload: Project[] }>(),
)

export const GetProjectsAction = createAction(
    "[project] get list",
)
