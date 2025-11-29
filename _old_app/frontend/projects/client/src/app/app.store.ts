import { projectsReducer } from "@client/entities/project"
import { tasksReducer } from "@client/entities/task"
import { userReducer } from "@client/entities/user"
import { statusReducer } from "@client/entities/status"
import { commentsReducer } from "@client/entities/comment"

import * as projectEffects from "@client/entities/project/model/project.effects"
import * as taskEffects from "@client/entities/task/model/task.effects"
import * as appEffects from "@client/shared/store/app.effects"
import * as userEffects from "@client/entities/user/model/user.effects"
import * as statusEffects from "@client/entities/status/model/status.effects"
import * as commentEffects from "@client/entities/comment/model/comment.effects"

export const reducers = {
	project: projectsReducer,
	task: tasksReducer,
	user: userReducer,
	status: statusReducer,
	comments: commentsReducer,
}

export const effects = [appEffects, projectEffects, taskEffects, userEffects, statusEffects, commentEffects]
