import { Component, inject, Input, OnInit } from "@angular/core"
import { FormControl, FormGroup, ReactiveFormsModule, Validators } from "@angular/forms"
import { Comment, CommentCreateAction } from "@client/entities/comment"
import { Task } from "@client/entities/task"
import { selectUser } from "@client/entities/user"
import { AppState } from "@client/shared/store"
import { Store } from "@ngrx/store"
import { filter, map, Observable } from "rxjs"

@Component({
	selector: "lu-task-comment-form",
	imports: [ReactiveFormsModule],
	template: `
		<form [formGroup]="form" (ngSubmit)="save()">
			<div class="mb-4">
				<label for="message">Message</label>
				<input id="message" type="text" formControlName="message" class="input" />
			</div>
			<button class="btn btn-primary" [disabled]="!form.valid">Submit</button>
		</form>
	`,
})
export class TaskCommentFormComponent implements OnInit {
	@Input() task: Task

	form: FormGroup = new FormGroup({
		message: new FormControl("", Validators.required),
	})

	private userId: number
	private store: Store<AppState> = inject(Store<AppState>)

	ngOnInit(): void {
		this.store.select(selectUser).pipe(
			filter(user => !!user),
			map(user => (this.userId = user.id))
		)
	}

	save(): void {
		const comment: Comment = {
			id: 0,
			entity_id: this.task.id,
			entity_type: "task",
			user_id: this.userId,
			content: this.form.value["message"],
			created_at: new Date(),
		}
		this.store.dispatch(CommentCreateAction({ comment }))
		this.form.reset()
	}
}
