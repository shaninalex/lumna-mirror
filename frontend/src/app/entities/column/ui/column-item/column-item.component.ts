import { AsyncPipe } from '@angular/common';
import type { OnInit } from '@angular/core';
import { Component, inject, Input } from '@angular/core';
import type { ColumnModel } from '@entities/column/model';
import type { TaskModel } from '@entities/task';
import { selectTasksByStatusId, TaskCardComponent, TaskInlineForm } from '@entities/task';
import { Store } from '@ngrx/store';
import type { Observable } from 'rxjs';

@Component({
    selector: 'lu-column-item',
    imports: [TaskCardComponent, TaskInlineForm, AsyncPipe],
    template: `
        <div class="card h-100">
            <div class="card-header d-flex justify-content-between align-items-center">
                <strong>{{ column.title }}</strong>
                @if (tasks$ | async; as tasks) {
                    <span class="badge text-bg-secondary">{{ tasks.length }}</span>
                }
            </div>
            <div class="card-body d-flex flex-column gap-3">
                @if (tasks$ | async; as tasks) {
                    @for (task of tasks; track $index) {
                        <lu-task-card [task]="task" />
                    }
                }

                <lu-task-inline-form
                    [column_id]="column.id"
                    [task_count]="0"
                />
            </div>
        </div>
    `,
})
export class ColumnItemComponent implements OnInit {
    private store = inject(Store);

    @Input({ required: true }) column: ColumnModel;
    tasks$: Observable<TaskModel[]>;

    ngOnInit(): void {
        this.tasks$ = this.store.select(selectTasksByStatusId(this.column.id));
    }
}
