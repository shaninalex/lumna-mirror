import { Component, computed, effect, inject, Input, signal } from '@angular/core';
import { form, required, Field } from '@angular/forms/signals';
import { projectEvents, ProjectStore } from '@entities/project';
import { Dispatcher, Events } from '@ngrx/signals/events';
import { ProjectEditModel } from './model/project-edit.model';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import {Title} from '@angular/platform-browser';

@Component({
    selector: 'app-project-edit-feature',
    imports: [Field],
    template: `
        <form (submit)="submit($event)">
            <div class="mb-2">
                <input class="border block w-full rounded-lg px-2 py-1 border-gray-400" placeholder="Project name" [field]="projectForm.name" />
                @if (projectForm.name().touched()) {
                    <ul class="error-list">
                        @for (error of projectForm.name().errors(); track error) {
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
                <button type="submit" class="bg-teal-500 text-white rounded-lg px-4 py-2 cursor-pointer">
                    @if(loading()) {
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
    @Input() projectId: number;
    private readonly store = inject(ProjectStore);
    private readonly events = inject(Events)
    loading = signal(false)
    project = computed(() => this.store.entities().find(p => p.id === this.projectId));
    private titleService = inject(Title);


    projectFormModel = signal<ProjectEditModel>({
        name: '',
    });
    errors = signal<string[]>([]);
    readonly dispatcher = inject(Dispatcher)

    constructor() {
        effect(() => {
            const p = this.project();
            if (p) {
                this.projectFormModel.set({ name: p.name });
                this.titleService.setTitle(`Edit Project: ${p.name}`)
            }
        });

        this.events
            .on(projectEvents.updateProject)
            .pipe(takeUntilDestroyed())
            .subscribe(() => this.loading.set(false));

    }

    projectForm = form(this.projectFormModel, (schemaPath) => {
        required(schemaPath.name, {message: 'Name is required'});
    });

    submit(event: Event): void {
        event.preventDefault()
        this.loading.set(true)
        this.dispatcher.dispatch(projectEvents.patch({ id: this.projectId, data: this.projectFormModel()}))
    }
}
