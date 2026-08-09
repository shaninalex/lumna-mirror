import { createAction, props } from '@ngrx/store';
import { ProjectModel, ProjectCreateModel } from './project.model';
import { Error } from '@shared/models'

export const actionProjectList = createAction(
    '[Project] get list',
    props<{ workspace_id: number }>(),
);

export const actionProjectDelete = createAction(
    '[Project] delete',
    props<{ project_id: number }>(),
);

export const actionProjectDeleteSuccess = createAction(
    '[Project] deleted',
    props<{ project_id: number }>(),
);

export const actionProjectsSetList = createAction(
    '[Project] set list',
    props<{ projects: ProjectModel[] }>(),
);

export const actionProjectCreate = createAction(
    '[Project] Create',
    props<{ payload: ProjectCreateModel }>(),
);

export const actionProjectUpdate = createAction(
    '[Project] update',
    props<{ id: number; data: ProjectCreateModel }>(),
);

export const actionProjectUpsert = createAction(
    '[Project] add',
    props<{ project: ProjectModel }>(),
);


export const actionProjectCreateFailed = createAction(
    "[Project] create failed",
    props<{ errors: Error[] }>()
);