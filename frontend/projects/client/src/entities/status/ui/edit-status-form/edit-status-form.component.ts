import {Component, inject, Input, OnInit} from '@angular/core';
import {DeleteStatusAction, PatchStatusAction, Status} from '@client/entities/status';
import {FormControl, FormGroup, ReactiveFormsModule, Validators} from '@angular/forms';
import {AppState} from '@client/shared/store';
import {Store} from '@ngrx/store';

@Component({
    selector: 'fr-edit-status-form',
    imports: [
        ReactiveFormsModule
    ],
    template: `
        <form [formGroup]="form" (ngSubmit)="onSubmit()">
            <div class="mb-2">
                <input type="text" class="input" formControlName="title"/>
            </div>
            <button class="btn btn-primary" type="submit">Update</button>
        </form>
        <hr class="my-2">
        <button class="btn btn-accent mb-2" (click)="onDelete()">Delete</button>
        <p class="text-sm">Deleting status also delete all tasks in it</p>
    `
})
export class EditStatusFormComponent implements OnInit {
    @Input() status: Status;
    private store = inject(Store<AppState>);

    form: FormGroup = new FormGroup({
        title: new FormControl('', Validators.required)
    })

    ngOnInit() {
        this.form.setValue({title: this.status.title})
    }

    onSubmit(): void {
        this.store.dispatch(PatchStatusAction({
            payload: {
                title: this.form.value['title'],
            },
            projectId: this.status.project_id,
            statusId: this.status.id,
        }))
    }

    onDelete(): void {
        this.store.dispatch(DeleteStatusAction({
            projectId: this.status.project_id,
            statusId: this.status.id,
        }))
    }
}
