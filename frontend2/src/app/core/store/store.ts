import { routerReducer } from "@ngrx/router-store";

import { WorkspaceEffects, workspaceReducer } from "@entities/workspace";

export const effects = [
    // Core app effects
    // RouterEffects,

    // Entity effects
    WorkspaceEffects
];

export const reducers = {
    workspace: workspaceReducer,
    // user: userReducer,
    // session: sessionReducer,
    // board: boardReducer,
    // column: columnReducer,
    // task: taskReducer,
    // activity: activityReducer,

    // -----
    router: routerReducer
};
