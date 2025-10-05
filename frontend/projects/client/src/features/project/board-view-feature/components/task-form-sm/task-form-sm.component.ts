import {Component, inject, Input} from '@angular/core';
import {LoaderComponent} from '@client/shared/ui/loader';
import {FormControl, FormGroup, ReactiveFormsModule, Validators} from '@angular/forms';
import {CreateTaskAction, SetTaskAction} from '@client/entities/task';
import {Actions, ofType} from '@ngrx/effects';
import {Store} from '@ngrx/store';
import {AppState} from '@client/shared/store';
import {Project} from '@client/entities/project';
import {StatusColumn} from '@client/features/project/board-view-feature/board.model';

@Component({
    selector: 'fr-task-form-sm',
    imports: [
        LoaderComponent,
        ReactiveFormsModule,
    ],
    template: `
        @if (showForm) {
            <form [formGroup]="form" (ngSubmit)="submitForm()" class="mb-4 flex gap-2">
                <div class="flex-grow">
                    <input class="input" autofocus placeholder="Task title" type="text" formControlName="title">
                    @if (form.controls['title'].dirty && form.controls['title'].errors) {
                        @if (form.controls['title'].errors['required']) {
                            <div class="text-sm">This field is required</div>
                        }
                        @if (form.controls['title'].errors['pattern']) {
                            <div class="text-sm">Special characters! Only a-z, A-Z and 0-9 are available</div>
                        }
                    }
                </div>

                <div class="flex gap-2">
                    @if (loading) {
                        <ui-loader/>
                    } @else {
                        <button class="btn btn-primary" type="submit" [disabled]="loading || !form.valid">
                            +
                        </button>
                    }
                    <button class="btn btn-secondary" type="button" (click)="cancel()">
                        X
                    </button>
                </div>
            </form>
        } @else {
            <button class="btn btn-secondary" (click)="showForm = true">
                Create task
            </button>
        }
    `
})
export class TaskFormSmComponent {
    @Input() project: Project;
    @Input() column: StatusColumn;
    showForm: boolean = false;
    loading: boolean = false;
    form: FormGroup = new FormGroup({
        'title': new FormControl({value: '', disabled: this.loading}, [Validators.required]),
    });

    private action$ = inject(Actions);
    private store = inject(Store<AppState>);

    constructor() {
        this.action$.pipe(ofType(SetTaskAction)).subscribe(() => {
            this.loading = false
        })
    }

    submitForm(): void {
        this.loading = true
        this.store.dispatch(CreateTaskAction({
            projectId: this.project.id,
            payload: {
                title: this.form.value['title'],
                status_id: this.column.status.id,
            }
        }))
        this.form.reset();
    }

    cancel(): void {
        this.form.reset();
        this.showForm = false;
    }
}
