import { Injectable } from "@angular/core"
import { map, Observable } from "rxjs"
import { environment as env } from "@client/environments/environment.development"
import { CommonApiService } from "@client/shared/common"
import { CreateTaskInput, Task, TaskDetailInput } from "@client/entities/task"

const tasksUrl = {
	tasksRoot: (projectId: number) => `${env.API_ROOT}/api/v1/project/${projectId}/tasks`,
	task: (taskId: number) => `${env.API_ROOT}/api/v1/task/${taskId}`,
}

@Injectable({ providedIn: "root" }) // TODO: does it has to be root?
export class TaskService extends CommonApiService {
	public List(projectId: number): Observable<Task[]> {
		return this.get<Task[]>(tasksUrl.tasksRoot(projectId)).pipe(map(data => data.data))
	}

	public Create(projectId: number, payload: CreateTaskInput): Observable<Task> {
		return this.post<Task>(tasksUrl.tasksRoot(projectId), payload).pipe(map(data => data.data))
	}

	public Detail(taskId: number): Observable<Task> {
		return this.get<Task>(tasksUrl.task(taskId)).pipe(map(data => data.data))
	}

	public Patch(taskId: number, payload: TaskDetailInput): Observable<Task> {
		return this.patch<Task>(tasksUrl.task(taskId), payload).pipe(map(data => data.data))
	}

	public Delete(taskId: number): Observable<null> {
		return this.delete<null>(tasksUrl.task(taskId)).pipe(map(data => data.data))
	}
}
