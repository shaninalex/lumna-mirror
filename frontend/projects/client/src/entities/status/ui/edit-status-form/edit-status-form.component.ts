import {Component, inject, OnInit} from '@angular/core';
import {DeleteStatusAction, PatchStatusAction} from '@client/entities/status';
import {FormControl, FormGroup, ReactiveFormsModule, Validators} from '@angular/forms';
import {AppState} from '@client/shared/store';
import {Store} from '@ngrx/store';

import {DialogData} from './model';
import {DialogRef, DIALOG_DATA} from '@angular/cdk/dialog';


@Component({
    selector: 'fr-edit-status-form',
    imports: [
        ReactiveFormsModule,
    ],
    template: `
        <div class="card">
            <div class="card-title mb-4">Config "{{ form.value['title'] }}" column</div>
            <div class="mb-4">
                <form [formGroup]="form">
                    <div class="mb-4">
                        <input class="input" formControlName="title"/>
                    </div>
                    <div>
                        <input id="task-complete" class="me-2" type="checkbox" formControlName="complete"/>
                        <label for="task-complete">Task complete</label>
                        <p class="text-sm">Tasks in that column will be marked as completed</p>
                    </div>
                </form>
            </div>

            <div class="flex gap-2">
                <button class="btn btn-secondary" (click)="cancel()">Cancel</button>
                <button class="btn btn-primary" (click)="onSubmit()">Update</button>
                <button class="btn btn-danger" (click)="onDelete()" title="Deleting status also delete all tasks in it">
                    Delete
                </button>
            </div>
        </div>
    `
})
export class EditStatusFormComponent implements OnInit {
    readonly dialogRef = inject(DialogRef<EditStatusFormComponent>);
    readonly data = inject<DialogData>(DIALOG_DATA);
    private store = inject(Store<AppState>);

    form: FormGroup = new FormGroup({
        title: new FormControl('', Validators.required),
        complete: new FormControl(false, Validators.required),
    })

    ngOnInit() {
        this.form.setValue({
            title: this.data.status.title,
            complete: this.data.status.complete,
        })
    }

    onSubmit(): void {
        this.store.dispatch(PatchStatusAction({
            payload: {
                title: this.form.value['title'],
                complete: this.form.value['complete'],
            },
            projectId: this.data.status.project_id,
            statusId: this.data.status.id,
        }))
        this.dialogRef.close(); //this.data.status
    }

    onDelete(): void {
        this.store.dispatch(DeleteStatusAction({
            projectId: this.data.status.project_id,
            statusId: this.data.status.id,
        }))
        this.dialogRef.close();
    }

    cancel(): void {
        this.dialogRef.close();
    }
}
