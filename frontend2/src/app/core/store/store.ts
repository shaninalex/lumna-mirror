import { routerReducer } from "@ngrx/router-store";

import { WorkspaceEffects, workspaceReducer } from "@entities/workspace";
import { TaskEffects, taskReducer } from "@entities/task";
import { ProjectEffects, projectReducer } from "@entities/project";
import { ListEffects, listReducer } from "@entities/list";

export const effects = [
    // Core app effects
    // RouterEffects,

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
    router: routerReducer
};
