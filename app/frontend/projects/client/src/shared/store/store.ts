import { ProjectState } from "@client/entities/project"
import { TasksState } from "@client/entities/task"
import { UserState } from "@client/entities/user"
import { StatusState } from "@client/entities/status"
import { CommentsState } from "@client/entities/comment"

export interface AppState {
	project: ProjectState
	task: TasksState
	user: UserState
	status: StatusState
	comments: CommentsState
}
