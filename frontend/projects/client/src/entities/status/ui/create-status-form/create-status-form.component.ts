import {Component, inject, Input} from '@angular/core';
import {FormControl, FormGroup, ReactiveFormsModule, Validators} from '@angular/forms';
import {Store} from '@ngrx/store';
import {AppState} from '@client/shared/store';
import {CreateStatusAction} from '@client/entities/status';
import {MatInputModule} from '@angular/material/input';
import {MatButtonModule} from '@angular/material/button';

@Component({
    selector: 'fr-create-status-form',
    imports: [
        ReactiveFormsModule,
        MatInputModule,
        MatButtonModule
    ],
    template: `
        <form [formGroup]="form" (ngSubmit)="onSubmit()" class="my-4">
            <mat-form-field appearance="outline" class="w-full">
                <input matInput type="text" placeholder="Status title" formControlName="title" />
            </mat-form-field>
            <button matButton="outlined" type="submit" [disabled]="!form.valid">Save</button>
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
