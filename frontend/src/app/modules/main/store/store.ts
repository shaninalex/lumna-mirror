import { ProjectEffects, projectReducer } from "@entities/project";
import { UserEffects, userReducer } from "@entities/user";
import { WorkspaceEffects, workspaceReducer } from "@entities/workspace";

export const mainReducers = {
    user: userReducer,
    workspace: workspaceReducer,
    project: projectReducer,
};

export const mainEffects = [
    UserEffects,
    WorkspaceEffects,
    ProjectEffects,
];