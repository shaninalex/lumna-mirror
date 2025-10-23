import {Component, ElementRef, HostListener, inject, Input, OnInit} from '@angular/core';
import {LoaderComponent} from '@client/shared/ui/loader';
import {FormControl, FormGroup, ReactiveFormsModule, Validators} from '@angular/forms';
import {TaskCreateAction, TaskSetAction} from '@client/entities/task';
import {Actions, ofType} from '@ngrx/effects';
import {Store} from '@ngrx/store';
import {AppState} from '@client/shared/store';
import {Project} from '@client/entities/project';
import {StatusColumn} from '@client/features/project/board-view-feature/board.model';

@Component({
    selector: 'lu-task-form-sm',
    imports: [
        LoaderComponent,
        ReactiveFormsModule,
    ],
    template: `
        @if (showForm) {
            <form [formGroup]="form" (ngSubmit)="submitForm()">
                <div class="mb-2">
                    <input class="input" autofocus placeholder="Task title" type="text" formControlName="title">
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
                    @if (loading) {
                        <ui-loader/>
                    } @else {
                        <button class="btn btn-primary btn-icon" type="submit" [disabled]="loading || !form.valid">
                            <i class="i-plus-circle text-xl"></i>
                        </button>
                    }
                    <button type="button" (click)="cancel()">
                        <i class="i-close-circle text-xl"></i>
                    </button>
                </div>
            </form>
        } @else {
            <button (click)="showForm = true" class="flex gap-2 items-center hover-space w-full">
                <i class="i-plus-circle text-xl"></i>
                Add a card
            </button>
        }
    `
})
export class TaskFormSmComponent implements OnInit {
    @Input() project: Project;
    @Input() column: StatusColumn;

    eRef = inject(ElementRef)

    showForm: boolean = false;
    loading: boolean = false;
    form: FormGroup = new FormGroup({
        'title': new FormControl({value: '', disabled: this.loading}, [Validators.required]),
    });

    private action$ = inject(Actions);
    private store = inject(Store<AppState>);

    ngOnInit() {
        this.action$.pipe(ofType(TaskSetAction)).subscribe(() => {
            this.loading = false
        })
    }

    @HostListener('document:click', ['$event'])
    clickout(event: { target: any; }) {
        if (!this.eRef.nativeElement.contains(event.target) && this.showForm ) {
            this.showForm = false
        }
    }

    submitForm(): void {
        this.loading = true
        this.store.dispatch(TaskCreateAction({
            projectId: this.project.id,
            payload: {
                title: this.form.value['title'],
                status_id: this.column.status.id,
            }
        }))
        this.form.reset();
    }

    cancel(): void {
        this.form.reset();
        this.showForm = false;
    }
}
