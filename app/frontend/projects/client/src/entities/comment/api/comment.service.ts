import { Injectable } from "@angular/core"
import { map, Observable } from "rxjs"
import { environment as env } from "@client/environments/environment.development"
import { CommonApiService } from "@client/shared/common"
import { Comment } from "../model"
import { HttpParams } from "@angular/common/http"

const tasksUrl = {
	comments: () => `${env.API_ROOT}/api/v1/comments`,
	commentDelete: (commentId: number) => `${env.API_ROOT}/api/v1/comments/${commentId}`,
}

@Injectable({ providedIn: "root" })
export class CommentService extends CommonApiService {
	public Create(payload: Comment): Observable<Comment> {
		return this.post<Comment>(tasksUrl.comments(), payload).pipe(map(data => data.data))
	}

	public List(entity_id: number, entity_type: string): Observable<Comment[]> {
		const params = new HttpParams().set("entity_id", entity_id).set("entity_type", entity_type)
		return this.get<Comment[]>(tasksUrl.comments(), params).pipe(map(data => data.data))
	}
}
