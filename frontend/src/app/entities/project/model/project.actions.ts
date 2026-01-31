import { createAction, props } from '@ngrx/store';
import { ProjectModel, ProjectPayload } from './project.model';

export const actionProjectList = createAction('[Project] get list');

export const actionProjectDelete = createAction(
    '[Project] delete',
    props<{ project_id: string }>(),
);

export const actionProjectDeleteSuccess = createAction('[Project] deleted');

export const actionProjectsSetList = createAction(
    '[Project] set list',
    props<{ projects: ProjectModel[] }>(),
);

export const actionProjectCreate = createAction(
    '[Project] Create',
    props<{ payload: ProjectPayload }>(),
);

export const actionProjectAdd = createAction('[Project] add', props<{ project: ProjectModel }>());
