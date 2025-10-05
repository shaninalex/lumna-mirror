import {Component, inject, Input} from '@angular/core';
import {FormControl, FormGroup, ReactiveFormsModule, Validators} from '@angular/forms';
import {Store} from '@ngrx/store';
import {AppState} from '@client/shared/store';
import {CreateStatusAction} from '@client/entities/status';

@Component({
    selector: 'fr-create-status-form',
    imports: [
        ReactiveFormsModule,
    ],
    template: `
        @if (showForm) {
            <form [formGroup]="form" (ngSubmit)="onSubmit()" class="my-4">
                <div class="w-full mb-4">
                    <input class="input" type="text" placeholder="Column title" formControlName="title"/>
                </div>
                <div class="mb-4">
                    <input id="status-completed" class="me-2" type="checkbox" formControlName="complete"/>
                    <label for="status-completed">Completed</label>
                    <p class="text-sm">Tasks in that column will be marked as completed</p>
                </div>
                <div class="flex gap-2">
                    <button class="btn" type="submit" [disabled]="!form.valid">Save</button>
                    <button class="btn-secondary" type="button" (click)="cancel()">Cancel</button>
                </div>
            </form>
        } @else {
            <button type="button" (click)="showForm = true">New column</button>
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
