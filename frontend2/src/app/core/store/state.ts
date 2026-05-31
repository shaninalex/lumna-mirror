import { AppState } from "@core/store/app";
import { SessionState } from "@root/src/app/core";
import { UserState } from "@entities/user";
import { WorkspaceState } from "@entities/workspace/model/workspace.store";
import { ProjectState } from "@entities/project";
import { ListState } from "@entities/list";
import { TaskState } from "vitest";

export interface ApplicationState {
    application: AppState;
    session: SessionState;
    user: UserState;
    workspace: WorkspaceState;
    project: ProjectState;
    list: ListState;
    task: TaskState;
}
