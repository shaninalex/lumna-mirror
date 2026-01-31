import { createAction, props } from '@ngrx/store';
import { ProjectModel, ProjectPayload } from './project.model';

export const actionProjectList = createAction('[Projects] get list');

export const actionProjectDelete = createAction(
    '[Projects] delete',
    props<{ project_id: string }>(),
);

export const actionProjectsSetList = createAction(
    '[Projects] set list',
    props<{ projects: ProjectModel[] }>(),
);

export const actionProjectCreate = createAction(
    '[Projects] Create',
    props<{ payload: ProjectPayload }>(),
);

export const actionProjectAdd = createAction('[Projects] add', props<{ project: ProjectModel }>());
