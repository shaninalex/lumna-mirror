import {Component, EventEmitter, inject, OnChanges, Output, SimpleChanges} from '@angular/core';
import {FormControl, FormGroup, ReactiveFormsModule, Validators} from '@angular/forms';
import {LoaderComponent} from '@client/shared/ui/loader';
import {Actions, ofType} from '@ngrx/effects';
import {CreateProjectAction, SetProjectAction} from '@client/entities/project';
import {Store} from '@ngrx/store';
import {AppState} from '@client/shared/store';

@Component({
    selector: "fr-new-project-form",
    imports: [
        ReactiveFormsModule,
        LoaderComponent,
    ],
    template: `
        <div>
            @if (!showForm) {
                <button class="btn btn-primary" (click)="toggleProjectForm()">Create Project</button>
            } @else {
                <form [formGroup]="form" (ngSubmit)="onSubmit()">
                    <div class="mb-2">
                        <input class="input" autofocus placeholder="Project Title" type="text" formControlName="title">
                        @if (form.controls['title'].dirty && form.controls['title'].errors) {
                            @if (form.controls['title'].errors['required']) {
                                <div class="text-sm">This field is required</div>
                            }
                            @if (form.controls['title'].errors['pattern']) {
                                <div class="text-sm">Special characters! Only a-z, A-Z and 0-9 are available</div>
                            }
                        }
                    </div>
                    <div class="flex gap-2 items-center">
                        <button class="btn btn-primary" [disabled]="loading || !form.valid" type="submit">Create</button>
                        <button class="btn btn-secondary" [disabled]="loading" type="button" (click)="cancel()">Cancel</button>
                        @if (loading) {
                            <ui-loader/>
                        }
                    </div>
                </form>
            }
        </div>
    `
})
export class NewProjectFormComponent implements OnChanges {
    @Output() onCancel: EventEmitter<boolean> = new EventEmitter<boolean>();
    loading: boolean = false;
    showForm: boolean = false;
    private actions$ = inject(Actions);
    private store = inject(Store<AppState>);

    form: FormGroup = new FormGroup({
        'title': new FormControl({value: '', disabled: this.loading}, [Validators.required]),
    })

    constructor() {
        this.actions$.pipe(ofType(SetProjectAction)).subscribe(() => {
            this.loading = this.showForm = false
        })
    }

    ngOnChanges(changes: SimpleChanges) {
        this.loading = changes["loading"].currentValue
        const control = this.form.get("title");
        if (this.loading) {
            control?.disable();
        } else {
            control?.enable();
        }
    }

    cancel(): void {
        this.showForm = false;
    }

    toggleProjectForm(): void {
        this.showForm = !this.showForm;
    }

    onSubmit(): void {
        this.loading = true;
        const project: Record<string, string> = { title: this.form.value['title'] }
        this.store.dispatch(CreateProjectAction({payload: project}))
    }
}
