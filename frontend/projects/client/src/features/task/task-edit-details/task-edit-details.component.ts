import {Component, EventEmitter, inject, Input, OnInit, Output} from '@angular/core';
import {MatFormFieldModule} from '@angular/material/form-field';
import {MatButtonModule} from '@angular/material/button';
import {MatInputModule} from '@angular/material/input';
import {FormControl, FormGroup, ReactiveFormsModule, Validators} from '@angular/forms';
import {Task, TaskService} from '@client/entities/task';

@Component({
    selector: 'ts-task-edit-details',
    imports: [
        MatFormFieldModule,
        MatInputModule,
        MatButtonModule,
        ReactiveFormsModule,
    ],
    template: `
        <form [formGroup]="form" (ngSubmit)="onSubmit()" class="flex flex-col gap-4">
            <mat-form-field appearance="outline">
                <mat-label>Clearable input</mat-label>
                <input matInput type="text" formControlName="title">
            </mat-form-field>
            <mat-form-field appearance="outline">
                <mat-label>Clearable input</mat-label>
                <textarea matInput type="text" formControlName="description"></textarea>
            </mat-form-field>
            <div>
                <button type="submit" matButton="outlined">Save</button>
            </div>
        </form>
    `
})
export class TaskEditDetailsComponent implements OnInit {
    @Input() task: Task;
    @Output() update: EventEmitter<any> = new EventEmitter<any>();

    form = new FormGroup({
        "title": new FormControl("", Validators.required),
        "description": new FormControl(""),
    })

    onSubmit(): void {
        this.update.emit(this.form.value)
    }

    ngOnInit() {
        this.form.setValue({
            title: this.task.title,
            description: this.task.description,
        })
    }
}
