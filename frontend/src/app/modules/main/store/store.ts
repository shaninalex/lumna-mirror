import { ProjectEffects, projectReducer } from "@entities/project";
import { UserEffects, userReducer } from "@entities/user";
import { WorkspaceEffects, workspaceReducer } from "@entities/workspace";
import { MainEffects } from "./main.effects";

export const mainReducers = {
    user: userReducer,
    workspace: workspaceReducer,
    project: projectReducer,
};

export const mainEffects = [
    MainEffects,
    UserEffects,
    WorkspaceEffects,
    ProjectEffects,
];