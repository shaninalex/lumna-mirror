import {createAction, props} from '@ngrx/store';
import {Project} from '@client/entities/project/model/project.model';

export const SetProjectsAction = createAction(
    "[project] set list",
    props<{ payload: Project[] }>(),
)

export const GetProjectsAction = createAction(
    "[project] get list",
)

export const SetProjectAction = createAction(
    "[project] set",
    props<{ payload: Project }>(),
)

export const CreateProjectAction = createAction(
    "[project] create",
    props<{ payload: Record<string, string> }>(),
)
