import { Task } from "@client/entities/task"
import { Status } from "@client/entities/status"

export interface StatusColumn {
	id: string
	title: string
	tasks: Task[]
	status: Status
}
