import {Component, inject, Input} from '@angular/core';
import {FormControl, FormGroup, ReactiveFormsModule, Validators} from '@angular/forms';
import {Store} from '@ngrx/store';
import {AppState} from '@client/shared/store';
import {CreateStatusAction} from '@client/entities/status';

@Component({
    selector: 'fr-create-status-form',
    imports: [
        ReactiveFormsModule
    ],
    template: `
        <form [formGroup]="form" (ngSubmit)="onSubmit()">
            <fieldset class="fieldset">
                <legend class="fieldset-legend">Status name</legend>
                <input type="text" class="input" placeholder="Type here" formControlName="title" />
            </fieldset>
            <button class="btn btn-primary" type="submit" [disabled]="!form.valid">Save</button>
        </form>
    `
})
export class CreateStatusFormComponent {
    @Input() projectId: number;
    private store = inject(Store<AppState>);

    form: FormGroup = new FormGroup({
        title: new FormControl("", Validators.required),
    })

    onSubmit(): void {
        this.store.dispatch(CreateStatusAction({
            payload: {
                title: this.form.value['title'],
            },
            projectId: this.projectId,
        }))
        this.form.reset()
    }
}
