import { routerReducer } from "@ngrx/router-store";

import { WorkspaceEffects, workspaceReducer } from "@entities/workspace";
import { TaskEffects, taskReducer } from "@entities/task";
import { ProjectEffects, projectReducer } from "@entities/project";
import { ListEffects, listReducer } from "@entities/list";

import { SessionEffects, sessionReducer } from "./session";

export const effects = [
    // Core app effects
    // RouterEffects,
    SessionEffects,

    // Entity effects
    WorkspaceEffects,
    ProjectEffects,
    ListEffects,
    TaskEffects
];

export const reducers = {
    workspace: workspaceReducer,
    project: projectReducer,
    list: listReducer,
    task: taskReducer,

    // -----
    session: sessionReducer,
    router: routerReducer
};
