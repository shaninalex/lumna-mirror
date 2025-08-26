import {Component, inject} from '@angular/core';
import {MAT_DIALOG_DATA, MatDialogContent} from '@angular/material/dialog';
import {Task} from '@client/entities/task';
import {MatInputModule} from '@angular/material/input';
import {FormControl, FormGroup, ReactiveFormsModule, Validators} from '@angular/forms';
import {MatButton} from '@angular/material/button';

export interface DialogData {
    task: Task;
}

@Component({
    selector: 'ts-task-detail-view',
    imports: [
        MatDialogContent,
        MatInputModule,
        ReactiveFormsModule,
        MatButton,
    ],
    template: `
        <mat-dialog-content class="w-screen max-w-7xl mx-auto">
            <form [formGroup]="form" class="flex flex-col gap-4">
                <mat-form-field appearance="outline">
                    <mat-label>Clearable input</mat-label>
                    <input matInput type="text" formControlName="title">
                </mat-form-field>

                <mat-form-field appearance="outline">
                    <mat-label>Clearable input</mat-label>
                    <input matInput type="text" formControlName="description">
                </mat-form-field>

                <div>
                    <button matButton="outlined">Save</button>
                </div>
            </form>
        </mat-dialog-content>`
})
export class TaskDetailViewComponent {
    data = inject<DialogData>(MAT_DIALOG_DATA);
    form = new FormGroup({
        "title": new FormControl(this.data.task.title, Validators.required),
        "description": new FormControl(this.data.task.description),
    })
}
