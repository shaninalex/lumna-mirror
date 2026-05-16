import { routerReducer } from "@ngrx/router-store";

import { WorkspaceEffects, workspaceReducer } from "@entities/workspace";
import { TaskEffects, taskReducer } from "@entities/task";
import { ProjectEffects, projectReducer } from "@entities/project";
import { ListEffects, listReducer } from "@entities/list";
import { UserEffects, userReducer } from "@entities/user";

import { SessionEffects, sessionReducer } from "./session";
import { ApplicationEffects } from "./app";
import { appReducer } from "./app/app.store";

export const effects = [
    // Core app effects
    // RouterEffects,
    SessionEffects,
    ApplicationEffects,

    // Entity effects
    UserEffects,
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
    // router: routerReducer,
    user: userReducer,
    application: appReducer
};
