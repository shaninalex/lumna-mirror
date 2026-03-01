import {Component, EventEmitter, Input, OnInit, Output, signal} from '@angular/core';
import {form, FormField, required} from '@angular/forms/signals';


@Component({
    selector: 'app-task-title-editor',
    imports: [
        FormField
    ],
    template: `
        <input
            class="form-control fw-bold form-control-lg"
            placeholder="Task title"
            [formField]="taskTitleForm.title"
            (blur)="onBlur()"
        />
        @if (taskTitleForm.title().touched() && taskTitleForm.title().errors().length) {
            @for (error of taskTitleForm.title().errors(); track error) {
                <div class="text-danger small">{{ error.message }}</div>
            }
        }
    `
})
export class TaskTitleEditorComponent implements OnInit {
    @Input() title: string;
    @Output() onChange: EventEmitter<string> = new EventEmitter();

    taskTitleFormModel = signal<{ title: string }>({ title: '' });
    taskTitleForm = form(this.taskTitleFormModel, (schemaPath) => {
        required(schemaPath.title, { message: 'Name is required' });
    });

    ngOnInit() {
        this.taskTitleFormModel.set({title: this.title});
    }

    onBlur(): void {
        if (this.taskTitleForm.title().errors().length) return
        this.onChange.emit(this.taskTitleForm.title().value())
    }
}
