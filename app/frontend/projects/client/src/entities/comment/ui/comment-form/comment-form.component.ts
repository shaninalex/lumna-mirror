import { Component, inject, Input, OnInit } from "@angular/core"
import { FormControl, FormGroup, ReactiveFormsModule, Validators } from "@angular/forms"
import { Comment, CommentCreateAction } from "@client/entities/comment"
import { selectUser } from "@client/entities/user"
import { AppState } from "@client/shared/store"
import { Store } from "@ngrx/store"
import { filter, map } from "rxjs"

@Component({
	selector: "lu-comment-form",
	imports: [ReactiveFormsModule],
	template: `
		<form [formGroup]="form" (ngSubmit)="save()">
			<div class="mb-4">
				<label for="message">Message</label>
				<input id="message" type="text" formControlName="message" class="input" placeholder="Enter your comment" />
			</div>
			<button class="btn btn-primary" [disabled]="!form.valid">Comment</button>
		</form>
	`,
})
export class CommentFormComponent implements OnInit {
	@Input() entity_id: number
	@Input() entity_type: string

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
			entity_id: this.entity_id,
			entity_type: this.entity_type,
			user_id: this.userId,
			content: this.form.value["message"],
			created_at: new Date(),
		}
		this.store.dispatch(CommentCreateAction({ comment }))
		this.form.reset()
	}
}
