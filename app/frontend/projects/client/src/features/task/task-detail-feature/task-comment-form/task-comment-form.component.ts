import { Component, inject, Input } from "@angular/core"
import { FormControl, FormGroup, ReactiveFormsModule, Validators } from "@angular/forms"
import { Task, TaskCreateCommentAction } from "@client/entities/task"
import { AppState } from "@client/shared/store"
import { Store } from "@ngrx/store"

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
export class TaskCommentFormComponent {
	private store: Store<AppState> = inject(Store<AppState>)
	@Input() task: Task
	form: FormGroup = new FormGroup({
		message: new FormControl("", Validators.required),
	})

	save(): void {
		this.store.dispatch(
			TaskCreateCommentAction({
				taskId: this.task.id,
				message: this.form.value["message"],
			})
		)
		this.form.reset()
	}
}
