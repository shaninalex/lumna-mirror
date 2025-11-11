import { Component, inject, Input, OnInit } from "@angular/core"
import { Comment, CommentComponent, CommentsListAction, selectComments } from "@client/entities/comment"
import { Store } from "@ngrx/store"
import { AppState } from "@client/shared/store"
import { Observable } from "rxjs"
import { AsyncPipe } from "@angular/common"

@Component({
	selector: "lu-comments-list",
	imports: [CommentComponent, AsyncPipe],
	template: `
		@if (comments | async; as comments) {
			<div class="flex flex-col gap-4">
				@for (comment of comments; track $index) {
					<lu-comment [comment]="comment" />
				}
			</div>
		}
	`,
})
export class CommentsListComponent implements OnInit {
	@Input() entity_id: number
	@Input() entity_type: string

	private store: Store<AppState> = inject(Store<AppState>)

	comments: Observable<Comment[]>

	ngOnInit(): void {
		this.store.dispatch(
			CommentsListAction({
				entity_id: this.entity_id,
				entity_type: "task",
			})
		)
		this.comments = this.store.select(selectComments(this.entity_id, this.entity_type))
	}
}
