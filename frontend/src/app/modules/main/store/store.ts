import { ProjectEffects } from "@entities/project";
import { UserEffects } from "@entities/user";
import { WorkspaceEffects } from "@entities/workspace";
import { MainEffects } from "./main.effects";
import { TaskEffects } from "@entities/task";
import { BoardsEffects } from "@entities/board";

export const mainEffects = [
    TaskEffects,
    MainEffects,
    UserEffects,
    WorkspaceEffects,
    ProjectEffects,
    BoardsEffects,
];