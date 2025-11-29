import { Injectable } from "@angular/core"
import { map, Observable } from "rxjs"
import { environment as env } from "@client/environments/environment.development"
import { Project, ProjectPatch } from "@client/entities/project"
import { CommonApiService } from "@client/shared/common"

const projectUrl = {
	root: `${env.API_ROOT}/api/v1/projects/`,
	item: (projectId: number) => `${env.API_ROOT}/api/v1/project/${projectId}`,
}

@Injectable({ providedIn: "root" }) // TODO: does it has to be root?
export class ProjectService extends CommonApiService {
	public List(): Observable<Array<Project>> {
		return this.get<Array<Project>>(projectUrl.root).pipe(map(data => data.data))
	}

	public Create(payload: Record<string, string>): Observable<Project> {
		return this.post<Project>(projectUrl.root, payload).pipe(map(data => data.data))
	}

	public Patch(projectId: number, payload: ProjectPatch): Observable<Project> {
		return this.patch<Project>(projectUrl.item(projectId), payload).pipe(map(data => data.data))
	}
}
