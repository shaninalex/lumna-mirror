import { ProjectEffects } from "@entities/project";
import { UserEffects } from "@entities/user";
import { WorkspaceEffects } from "@entities/workspace";
import { MainEffects } from "./main.effects";
import { TaskEffects } from "@entities/task";
import { ListsEffects } from "@entities/list";
import { StatusEffects } from "@entities/status/model/status.effects";

export const mainEffects = [
    TaskEffects,
    MainEffects,
    UserEffects,
    WorkspaceEffects,
    ProjectEffects,
    ListsEffects,
    StatusEffects,
];