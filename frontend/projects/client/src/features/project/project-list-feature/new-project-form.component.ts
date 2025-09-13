import {Component, EventEmitter, Input, OnChanges, Output, SimpleChanges} from '@angular/core';
import {FormControl, FormGroup, ReactiveFormsModule, Validators} from '@angular/forms';
import {LoaderComponent} from '@client/shared/ui/loader';

@Component({
    selector: "fr-new-project-form",
    imports: [
        ReactiveFormsModule,
        LoaderComponent,
    ],
    template: `
        <form [formGroup]="form" (ngSubmit)="submitForm()">
            <div class="mb-4">
                <label for="projectTitle">Project Title</label>
                <input id="projectTitle" type="text" class="input" formControlName="title" pattern="[a-zA-Z0-9 ]*">
                @if (form.controls['title'].dirty && form.controls['title'].errors) {
                    @if (form.controls['title'].errors['required']) {
                        <small class="text-red-500">This field is required</small>
                    }
                    @if (form.controls['title'].errors['pattern']) {
                        <small class="text-red-500">Special characters! Only a-z, A-Z and 0-9 are available</small>
                    }
                }
            </div>

            <div class="flex gap-2 items-center">
                <button [disabled]="loading || !form.valid" class="btn" type="submit">Create</button>
                <button [disabled]="loading" class="btn btn-secondary" type="button" (click)="cancel()">Cancel</button>
                @if (loading) {
                    <ui-loader />
                }
            </div>
        </form>
    `
})
export class NewProjectFormComponent implements OnChanges {
    @Output() onSubmit: EventEmitter<string> = new EventEmitter<string>();
    @Output() onCancel: EventEmitter<boolean> = new EventEmitter<boolean>();
    @Input() loading: boolean = false;

    form: FormGroup = new FormGroup({
        'title': new FormControl({value: '', disabled: this.loading}, [Validators.required]),
    })

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
        this.onCancel.emit()
    }

    submitForm(): void {
        this.onSubmit.emit(this.form.value['title'])
    }
}
