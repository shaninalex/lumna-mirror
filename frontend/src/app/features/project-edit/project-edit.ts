import { Component, inject, Input, signal } from '@angular/core';
import { form, required, FormField } from '@angular/forms/signals';
import { ProjectModel, ProjectState } from '@entities/project';
import { Dispatcher, Events } from '@ngrx/signals/events';
import { ProjectEditModel } from './model/project-edit.model';
import { Store } from '@ngrx/store';

@Component({
    selector: 'app-project-edit-feature',
    imports: [FormField],
    template: `
        <form (submit)="submit($event)">
            <div class="mb-2">
                <input class="input" placeholder="Project name" [formField]="projectForm.title" />
                @if (projectForm.title().touched()) {
                    <ul class="error-list">
                        @for (error of projectForm.title().errors(); track error) {
                            <li class="text-red-500 text-sm">{{ error.message }}</li>
                        }
                    </ul>
                }
            </div>

            @if (errors()) {
                <ul class="mb-2">
                    @for (error of errors(); track error) {
                        <li class="text-red-500 text-sm">{{ error }}</li>
                    }
                </ul>
            }

            <div>
                <button type="submit" class="btn btn-primary">
                    @if (loading()) {
                        Processing...
                    } @else {
                        Save
                    }
                </button>
            </div>
        </form>
    `,
})
export class ProjectEditFeature {
    @Input() projectId: string;
    private readonly store = inject(Store<ProjectState>);
    private readonly events = inject(Events);
    loading = signal(false);
    project: ProjectModel;

    projectFormModel = signal<ProjectEditModel>({
        title: '',
    });
    errors = signal<string[]>([]);
    readonly dispatcher = inject(Dispatcher);

    constructor() {
        // effect(() => {
        //     const p = this.project();
        //     if (p) {
        //         this.projectFormModel.set({ title: p.title });
        //     }
        // });
        // this.events
        //     .on(projectEvents.updateProject)
        //     .pipe(takeUntilDestroyed())
        //     .subscribe(() => this.loading.set(false));
    }

    projectForm = form(this.projectFormModel, (schemaPath) => {
        required(schemaPath.title, { message: 'Title is required' });
    });

    submit(event: Event): void {
        // event.preventDefault();
        // this.loading.set(true);
        // this.dispatcher.dispatch(
        //     // projectEvents.patch({ id: this.projectId, data: this.projectFormModel() }),
        // );
    }
}
