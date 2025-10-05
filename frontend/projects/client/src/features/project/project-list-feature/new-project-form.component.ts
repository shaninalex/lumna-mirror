import {Component, EventEmitter, inject, Input, OnChanges, Output, SimpleChanges} from '@angular/core';
import {FormControl, FormGroup, ReactiveFormsModule, Validators} from '@angular/forms';
import {LoaderComponent} from '@client/shared/ui/loader';
import {MatButtonModule} from '@angular/material/button';
import {MatInputModule} from '@angular/material/input';
import {MatCardModule} from '@angular/material/card';
import {Actions, ofType} from '@ngrx/effects';
import {CreateProjectAction, SetProjectAction} from '@client/entities/project';
import {Store} from '@ngrx/store';
import {AppState} from '@client/shared/store';

@Component({
    selector: "fr-new-project-form",
    imports: [
        ReactiveFormsModule,
        LoaderComponent,
        MatButtonModule,
        MatInputModule,
        MatCardModule,
    ],
    template: `
        <mat-card appearance="outlined">
            <mat-card-content>
                @if (!showForm) {
                    <button matButton="outlined" (click)="toggleProjectForm()">Create Project</button>
                } @else {
                    <form [formGroup]="form" (ngSubmit)="onSubmit()">
                        <mat-form-field appearance="outline">
                            <input matInput autofocus placeholder="Project Title" type="text" formControlName="title">
                            @if (form.controls['title'].dirty && form.controls['title'].errors) {
                                @if (form.controls['title'].errors['required']) {
                                    <mat-error>This field is required</mat-error>
                                }
                                @if (form.controls['title'].errors['pattern']) {
                                    <mat-error>Special characters! Only a-z, A-Z and 0-9 are available</mat-error>
                                }
                            }
                        </mat-form-field>

                        <div class="flex gap-2 items-center">
                            <button matButton="outlined" [disabled]="loading || !form.valid" type="submit">Create
                            </button>
                            <button matButton="outlined" [disabled]="loading" type="button" (click)="cancel()">Cancel
                            </button>
                            @if (loading) {
                                <ui-loader/>
                            }
                        </div>
                    </form>
                }
            </mat-card-content>
        </mat-card>
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
