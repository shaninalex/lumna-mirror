import { provideEffects } from "@ngrx/effects";
import { provideState } from "@ngrx/store";
import { WorkspaceApi } from "./api";
import { WorkspaceEffects, workspaceReducer } from "./model";

export function provideWorkspaceFeature() {
    console.log("[PROVIDE] workspaces");
    return [
        provideState("workspace", workspaceReducer),
        provideEffects(WorkspaceEffects),

        WorkspaceApi
    ];
}
