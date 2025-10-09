import {Component, inject, Input} from '@angular/core';
import {LoaderComponent} from '@client/shared/ui/loader';
import {FormControl, FormGroup, ReactiveFormsModule, Validators} from '@angular/forms';
import {TaskCreateAction, TaskSetAction} from '@client/entities/task';
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
                        <button class="btn btn-primary btn-icon" type="submit" [disabled]="loading || !form.valid">
                            <i class="i-plus-circle text-lg"></i>
                        </button>
                    }
                    <button class="btn btn-secondary btn-icon" type="button" (click)="cancel()">
                        <i class="i-close-circle text-lg"></i>
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
        this.action$.pipe(ofType(TaskSetAction)).subscribe(() => {
            this.loading = false
        })
    }

    submitForm(): void {
        this.loading = true
        this.store.dispatch(TaskCreateAction({
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
