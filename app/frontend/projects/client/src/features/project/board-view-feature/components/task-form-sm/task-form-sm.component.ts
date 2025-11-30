import { AfterViewChecked, Component, ElementRef, HostListener, inject, Input, OnInit, ViewChild } from "@angular/core"
import { LoaderComponent } from "@client/shared/ui/loader"
import { FormControl, FormGroup, ReactiveFormsModule, Validators } from "@angular/forms"
import { TaskCreateAction, TaskSetAction } from "@client/entities/task"
import { Actions, ofType } from "@ngrx/effects"
import { Store } from "@ngrx/store"
import { AppState } from "@client/shared/store"
import { Project } from "@client/entities/project"
import { StatusColumn } from "@client/features/project/board-view-feature/board.model"

@Component({
	selector: "lu-task-form-sm",
	imports: [LoaderComponent, ReactiveFormsModule],
	template: `
		@if (showForm) {
			<form [formGroup]="form" (ngSubmit)="submitForm()">
				<div class="mb-2">
					<input
						#titleInput
						class="input"
						placeholder="Task title"
						type="text"
						formControlName="title"
					/>
					@if (form.controls["title"].dirty && form.controls["title"].errors) {
						@if (form.controls["title"].errors["required"]) {
							<div class="text-sm">This field is required</div>
						}
						@if (form.controls["title"].errors["pattern"]) {
							<div class="text-sm">Special characters! Only a-z, A-Z and 0-9 are available</div>
						}
					}
				</div>

				<div class="flex items-center gap-2">
					@if (loading) {
						<ui-loader />
					} @else {
						<button class="btn btn-primary btn-icon" type="submit" [disabled]="loading || !form.valid">
							<i class="i-plus-circle text-xl"></i>
						</button>
					}
					<button type="button" (click)="cancel()">
						<i class="i-close-circle text-xl"></i>
					</button>
				</div>
			</form>
		} @else {
			<button (click)="toggleForm()" class="hover-space flex w-full items-center gap-2">
				<i class="i-plus-circle text-xl"></i>
				Add a card
			</button>
		}
	`,
})
export class TaskFormSmComponent implements OnInit, AfterViewChecked {
	@Input() project!: Project
	@Input() column!: StatusColumn

	@ViewChild("titleInput") titleInput!: ElementRef<HTMLInputElement>

	eRef = inject(ElementRef)
	private action$ = inject(Actions)
	private store = inject(Store<AppState>)

	showForm = false
	loading = false
	form = new FormGroup({
		title: new FormControl("", [Validators.required]),
	})

	private shouldFocus = false

	ngOnInit() {
		this.action$.pipe(ofType(TaskSetAction)).subscribe(() => {
			this.loading = false
		})
	}

	ngAfterViewChecked() {
		if (this.shouldFocus && this.titleInput) {
			this.titleInput.nativeElement.focus()
			this.shouldFocus = false
		}
	}

	toggleForm() {
		this.showForm = true
		this.shouldFocus = true
	}

	@HostListener("document:click", ["$event"])
	clickout(event: MouseEvent) {
		if (!this.eRef.nativeElement.contains(event.target) && this.showForm) {
			this.showForm = false
		}
	}

	submitForm() {
		this.loading = true
		this.store.dispatch(
			TaskCreateAction({
				projectId: this.project.id,
				payload: {
					title: this.form.value["title"]!,
					status_id: this.column.status.id,
				},
			})
		)
		this.form.reset()
	}

	cancel() {
		this.form.reset()
		this.showForm = false
	}
}
