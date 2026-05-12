import { createAction, props } from "@ngrx/store";
import { WorkspaceModel } from "./workspace.model";

export const actionWorkspaceGetList = createAction("[Workspace] get list");

export const actionWorkspaceSetList = createAction(
    "[Workspace] set list",
    props<{ list: WorkspaceModel[] }>()
);
