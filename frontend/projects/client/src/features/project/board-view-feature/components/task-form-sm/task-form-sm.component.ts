import {Component, inject, Input} from '@angular/core';
import {LoaderComponent} from '@client/shared/ui/loader';
import {FormControl, FormGroup, ReactiveFormsModule, Validators} from '@angular/forms';
import {CreateTaskAction, SetTaskAction} from '@client/entities/task';
import {Actions, ofType} from '@ngrx/effects';
import {Store} from '@ngrx/store';
import {AppState} from '@client/shared/store';
import {Project} from '@client/entities/project';
import {StatusColumn} from '@client/features/project/board-view-feature/board.model';
import {MatButtonModule} from '@angular/material/button';
import {MatInputModule} from '@angular/material/input';
import {MatIconModule} from '@angular/material/icon';

@Component({
    selector: 'fr-task-form-sm',
    imports: [
        LoaderComponent,
        ReactiveFormsModule,
        MatButtonModule,
        MatInputModule,
        MatIconModule,
    ],
    template: `
        @if (showForm) {

        <form [formGroup]="form" (ngSubmit)="submitForm()" class="mb-4">
            <div class="flex-grow">
                <mat-form-field appearance="outline" class="w-full">
                    <input matInput placeholder="Task title" type="text" formControlName="title">
                    @if (form.controls['title'].dirty && form.controls['title'].errors) {
                        @if (form.controls['title'].errors['required']) {
                            <mat-error class="text-sm">This field is required</mat-error>
                        }
                        @if (form.controls['title'].errors['pattern']) {
                            <mat-error class="text-sm">Special characters! Only a-z, A-Z and 0-9 are available</mat-error>
                        }
                    }
                </mat-form-field>
            </div>
            @if (loading) {
                <ui-loader/>
            } @else {
                <button matButton="outlined" type="submit"
                    [disabled]="loading || !form.valid" >
                    <mat-icon>add_circle</mat-icon>
                    Add
                </button>
            }
            <button matButton="outlined" type="button" (click)="cancel()">
                Cancel
            </button>
        </form>
        } @else {
            <button matButton="outlined" (click)="showForm = true">
                <mat-icon>add_circle</mat-icon>
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
            payload: {
                title: this.form.value['title'],
                status_id: this.column.id,
                project_code: this.project.code,
            }
        }))
    }

    cancel(): void {
        this.form.reset();
        this.showForm = false;
    }
}
