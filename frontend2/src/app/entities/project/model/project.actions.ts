import { createAction, props } from "@ngrx/store";
import { ProjectCreateModel, ProjectModel } from "./project.model";
import { Error } from "@shared/models";

export const actionProjectGetList = createAction("[Project] get list");

export const actionProjectSetList = createAction(
    "[Project] set list",
    props<{ list: ProjectModel[] }>()
);

export const actionProjectCreate = createAction(
    "[Project] create",
    props<{ data: ProjectCreateModel }>()
);

export const actionProjectCreateFailed = createAction(
    "[Project] create failed",
    props<{ errors: Error[] }>()
);

export const actionProjectCreateSuccessful = createAction(
    "[Project] create successful",
    props<{ project: ProjectModel }>()
);

export const actionProjectSetCurrentProjectId = createAction(
    "[Project] set current project_id",
    props<{ project_id: number }>()
);
