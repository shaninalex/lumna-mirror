import {Component, EventEmitter, Input, Output} from '@angular/core';
import {Project} from '@client/entities/project';
import {StatusColumn} from '@client/features/project/board-view-feature/board.model';
import {FormsModule, ReactiveFormsModule} from '@angular/forms';
import {TaskFormSmComponent} from '../task-form-sm';
import {EditStatusFormComponent} from '@client/entities/status';

@Component({
    selector: 'fr-column-header',
    template: `
        <div class="mb-2 flex justify-between">
            <div class="font-bold text-gray-500 dark:text-gray-300">{{ column.title }}</div>
            <div class="flex justify-between gap-2">
                <button (click)="toggleMenu()">...</button>
                <button (click)="toggleForm()" class="text-xl font-bold text-gray-500 dark:text-gray-300 cursor-pointer">+</button>
            </div>
        </div>
        @if (newTaskForm) {
            <fr-task-form-sm [project]="project" [column]="column" />
        }
        @if (showStatusMenu) {
            <fr-edit-status-form [status]="column.status" />
        }
    `,
    imports: [
        ReactiveFormsModule,
        FormsModule,
        TaskFormSmComponent,
        EditStatusFormComponent
    ]
})
export class ColumnHeaderComponent {
    @Input() project: Project;
    @Input() column: StatusColumn;
    newTaskForm: boolean = false;
    showStatusMenu: boolean = false;

    toggleForm(): void {
        this.newTaskForm = !this.newTaskForm
    }

    toggleMenu(): void {
        this.showStatusMenu = !this.showStatusMenu
    }
}
