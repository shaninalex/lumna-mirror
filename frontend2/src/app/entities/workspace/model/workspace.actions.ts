import { createAction, props } from "@ngrx/store";
import { WorkspaceCreateModel, WorkspaceModel } from "./workspace.model";
import { Error } from "@shared/models";

export const actionWorkspaceGetList = createAction("[Workspace] get list");

export const actionWorkspaceSetList = createAction(
    "[Workspace] set list",
    props<{ list: WorkspaceModel[] }>()
);

export const actionWorkspaceCreate = createAction(
    "[Workspace] create",
    props<{ data: WorkspaceCreateModel }>()
);

export const actionWorkspaceCreateSuccess = createAction(
    "[Workspace] create",
    props<{ data: WorkspaceModel }>()
);

export const actionWorkspaceCreateFailed = createAction(
    "[Workspace] create",
    props<{ errors: Error[] }>()
);
