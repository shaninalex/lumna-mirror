import { provideEffects } from "@ngrx/effects";
import { provideState } from "@ngrx/store";
import { WorkspaceApi } from "./api";
import { WorkspaceEffects, workspaceReducer } from "./model";
import { isDevMode } from "@angular/core";

export function provideWorkspaceFeature() {
    if (isDevMode()) {
        console.info("[PROVIDE] workspaces");
    }

    return [
        provideState("workspace", workspaceReducer),
        provideEffects(WorkspaceEffects),

        WorkspaceApi
    ];
}
