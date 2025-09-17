import {Component, EventEmitter, Input, Output} from '@angular/core';
import {Project} from '@client/entities/project';
import {StatusColumn} from '@client/features/project/board-view-feature/board.model';
import {FormsModule, ReactiveFormsModule} from '@angular/forms';
import {TaskFormSmComponent} from '../task-form-sm';

@Component({
    selector: 'fr-column-header',
    template: `
        <div class="mb-2 flex justify-between">
            <div class="font-bold text-slate-500 dark:text-slate-300">{{ column.title }}</div>
            <button (click)="toggleForm()" class="text-xl font-bold text-slate-500 dark:text-slate-300 cursor-pointer">+</button>
        </div>
        @if (newTaskForm) {
            <fr-task-form-sm [project]="project" [column]="column" />
        }
    `,
    imports: [
        ReactiveFormsModule,
        FormsModule,
        TaskFormSmComponent
    ]
})
export class ColumnHeaderComponent {
    @Input() project: Project;
    @Input() column: StatusColumn;
    newTaskForm: boolean = false

    toggleForm(): void {
        this.newTaskForm = !this.newTaskForm
    }
}
