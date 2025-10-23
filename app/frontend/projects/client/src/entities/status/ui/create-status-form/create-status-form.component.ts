import { Component, inject, Input } from '@angular/core';
import { FormControl, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { Store } from '@ngrx/store';
import { AppState } from '@client/shared/store';
import { StatusCreateAction } from '@client/entities/status';

@Component({
    selector: 'lu-create-status-form',
    imports: [
        ReactiveFormsModule,
    ],
    template: `
        @if (showForm) {
           <div class="card w-xs">
                <form [formGroup]="form" (ngSubmit)="onSubmit()" class="my-4">
                    <div class="w-full mb-4">
                        <input class="input" type="text" placeholder="Column title" formControlName="title"/>
                    </div>
                    <div class="flex gap-2">
                        <button class="btn btn-primary" type="submit" [disabled]="!form.valid">Save</button>
                        <button class="text-2xl cursor-pointer" type="button" (click)="cancel()">
                            <i class="i-close-circle"></i>
                        </button>
                    </div>
                </form>
            </div>
        } @else {
            <button class="btn btn-icon btn-secondary text-gray-600" type="button" (click)="showForm = true">
                <i class="i-plus-circle text-lg"></i>
                New column
            </button>
        }
    `
})
export class CreateStatusFormComponent {
    @Input() projectId: number;
    private store = inject(Store<AppState>);
    showForm: boolean = false;

    form: FormGroup = new FormGroup({
        title: new FormControl("", Validators.required),
    })

    onSubmit(): void {
        this.store.dispatch(StatusCreateAction({
            payload: {
                title: this.form.value['title'],
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
