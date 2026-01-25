import { Component, inject, signal } from '@angular/core';
import { DialogRef } from '@angular/cdk/dialog';
import { ProjectPayload } from '@entities/project';
import { Field, form, required, validate } from '@angular/forms/signals';

@Component({
    selector: 'app-project-form',
    imports: [Field],
    templateUrl: './project-form.html',
    styleUrl: './project-form.css',
})
export class ProjectForm {
    dialogRef = inject<DialogRef<ProjectPayload>>(DialogRef<ProjectPayload>);

    projectFormModel = signal<ProjectPayload>({
        title: '',
    });

    projectForm = form(this.projectFormModel, (schemaPath) => {
        required(schemaPath.title, { message: 'Title is required' });
        validate(schemaPath.title, ({ value }) => {
            if (value().trim().length === 0) {
                return {
                    kind: 'string',
                    message: 'Title should not be empty string',
                };
            }
            return null;
        });
    });

    submit(event: Event): void {
        event.preventDefault();
        this.dialogRef.close(this.projectFormModel());
    }

    cancel(): void {
        this.dialogRef.close();
    }
}
