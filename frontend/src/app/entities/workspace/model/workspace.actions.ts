import { createAction, props } from "@ngrx/store";
import { Error } from '@shared/models'
import { WorkspaceCreateModel, WorkspaceModel } from "./workspace.model";

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
    "[Workspace] created",
    props<{ data: WorkspaceModel }>()
);

export const actionWorkspaceCreateFailed = createAction(
    "[Workspace] create failed",
    props<{ errors: Error[] }>()
);

export const actionWorkspaceListRequested = createAction(
    "[Workspace] data requested",
    props<{ initiator?: any }>()
);

export const actionWorkspaceSetCurrent = createAction(
    "[Workspace] set current",
    props<{ id: number | null }>()
);
