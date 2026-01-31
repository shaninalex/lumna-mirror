import { createAction, props } from '@ngrx/store';
import { ProjectModel } from './project.model';

export const actionProjectList = createAction('[Projects] Get List');
export const actionProjectDelete = createAction(
    '[Projects] Delete',
    props<{ project_id: string }>(),
);
export const actionProjectsSet = createAction(
    '[Projects] Set list',
    props<{ projects: ProjectModel[] }>(),
);
