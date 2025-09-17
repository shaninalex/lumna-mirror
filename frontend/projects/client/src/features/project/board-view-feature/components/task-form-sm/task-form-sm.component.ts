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
        ReactiveFormsModule
    ],
    template: `
        <form [formGroup]="form" (ngSubmit)="submitForm()" class="mb-4">
            <div class="flex item-center gap-1">
                <div class="flex-grow">
                    <input placeholder="Task title" type="text" class="input block" formControlName="title"
                           pattern="[a-zA-Z0-9 ]*">
                </div>
                <div class="flex items-center">
                    @if (loading) {
                        <ui-loader/>
                    } @else {
                        <button
                            class="btn"
                            [disabled]="loading || !form.valid" type="submit">
                            +
                        </button>
                    }
                </div>
            </div>
            @if (form.controls['title'].dirty && form.controls['title'].errors) {
                @if (form.controls['title'].errors['required']) {
                    <small class="text-red-500">This field is required</small>
                }
                @if (form.controls['title'].errors['pattern']) {
                    <small class="text-red-500">Special characters! Only a-z, A-Z and 0-9 are available</small>
                }
            }
        </form>
    `
})
export class TaskFormSmComponent {
    @Input() project: Project;
    @Input() column: StatusColumn;

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
            payload: {
                title: this.form.value['title'],
                status_id: this.column.id,
                project_code: this.project.project_key,
            }
        }))
    }
}
