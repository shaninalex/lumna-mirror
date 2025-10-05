import {Component, inject, Input} from '@angular/core';
import {FormControl, FormGroup, ReactiveFormsModule, Validators} from '@angular/forms';
import {Store} from '@ngrx/store';
import {AppState} from '@client/shared/store';
import {CreateStatusAction} from '@client/entities/status';
import {MatInputModule} from '@angular/material/input';
import {MatButtonModule} from '@angular/material/button';
import {MatCheckboxModule} from '@angular/material/checkbox';

@Component({
    selector: 'fr-create-status-form',
    imports: [
        ReactiveFormsModule,
        MatInputModule,
        MatButtonModule,
        MatCheckboxModule,
    ],
    template: `
        @if (showForm) {
            <form [formGroup]="form" (ngSubmit)="onSubmit()" class="my-4">
                <mat-form-field appearance="outline" class="w-full">
                    <input matInput type="text" placeholder="Column title" formControlName="title"/>
                </mat-form-field>
                <div class="mb-4">
                    <mat-checkbox formControlName="complete">Complete</mat-checkbox>
                    <p class="text-sm">Tasks in that column will be marked as completed</p>
                </div>
                <button matButton="outlined" type="submit" [disabled]="!form.valid">Save</button>
                <button matButton="outlined" type="button" (click)="cancel()">Cancel</button>
            </form>
        } @else {
            <button matButton="outlined" type="button" (click)="showForm = true">New column</button>
        }
    `
})
export class CreateStatusFormComponent {
    @Input() projectId: number;
    private store = inject(Store<AppState>);
    showForm: boolean = false;

    form: FormGroup = new FormGroup({
        title: new FormControl("", Validators.required),
        complete: new FormControl(false, Validators.required),
    })

    onSubmit(): void {
        this.store.dispatch(CreateStatusAction({
            payload: {
                title: this.form.value['title'],
                complete: this.form.value['complete'],
            },
            projectId: this.projectId,
        }))
        this.form.reset()
    }

    cancel(): void {
        this.showForm = false;
        this.form.reset();
    }
}
