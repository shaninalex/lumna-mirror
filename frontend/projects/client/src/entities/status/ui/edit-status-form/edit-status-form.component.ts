import {Component, inject, OnInit} from '@angular/core';
import {DeleteStatusAction, PatchStatusAction} from '@client/entities/status';
import {FormControl, FormGroup, ReactiveFormsModule, Validators} from '@angular/forms';
import {AppState} from '@client/shared/store';
import {Store} from '@ngrx/store';

import {MatButton} from '@angular/material/button';
import {
    MAT_DIALOG_DATA,
    MatDialogActions,
    MatDialogContent,
    MatDialogRef,
    MatDialogTitle,
} from '@angular/material/dialog';
import {MatFormField} from '@angular/material/form-field';
import {MatInput} from '@angular/material/input';
import {DialogData} from './model';
import {MatTooltipModule} from '@angular/material/tooltip';
import {MatCheckbox} from '@angular/material/checkbox';


@Component({
    selector: 'fr-edit-status-form',
    imports: [
        ReactiveFormsModule,
        MatButton,
        MatDialogActions,
        MatDialogContent,
        MatDialogTitle,
        MatFormField,
        MatInput,
        MatTooltipModule,
        MatCheckbox
    ],
    template: `
        <h2 mat-dialog-title>Config "{{ data.status.title }}" column</h2>

        <mat-dialog-content>
            <form [formGroup]="form">
                <mat-form-field appearance="outline" class="w-full">
                    <input matInput formControlName="title" />
                </mat-form-field>
                <div>
                    <mat-checkbox formControlName="complete">Complete</mat-checkbox>
                    <p class="text-sm">Tasks in that column will be marked as completed</p>
                </div>
            </form>
        </mat-dialog-content>

        <mat-dialog-actions>
            <button matButton (click)="cancel()">Cancel</button>
            <button matButton (click)="onSubmit()">Update</button>
            <button matButton (click)="onDelete()"
                    matTooltip="Deleting status also delete all tasks in it">Delete</button>
        </mat-dialog-actions>
    `
})
export class EditStatusFormComponent implements OnInit {
    readonly dialogRef = inject(MatDialogRef<EditStatusFormComponent>);
    readonly data = inject<DialogData>(MAT_DIALOG_DATA);
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
        this.dialogRef.close(this.data.status);
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
