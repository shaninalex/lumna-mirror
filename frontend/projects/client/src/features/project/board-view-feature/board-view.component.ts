import {Component, inject, Input, OnInit} from '@angular/core';
import {
    TaskChangeStatusAction,
    TaskChangeStatusSuccessAction,
    selectTasksByProjectID,
    Task,
    TaskCardComponent
} from '@client/entities/task';
import {
    CdkDrag,
    CdkDragDrop,
    CdkDropList,
    CdkDropListGroup,
    moveItemInArray,
    transferArrayItem,
} from '@angular/cdk/drag-drop';
import {BoardViewApiService} from './api';
import {StatusColumn} from './board.model';
import {Store} from '@ngrx/store';
import {AppState} from '@client/shared/store';
import {Project} from '@client/entities/project';
import {ColumnHeaderComponent} from '@client/features/project/board-view-feature/components';
import {CreateStatusFormComponent, selectProjectStatusList} from '@client/entities/status';
import {combineLatest, map, Observable} from 'rxjs';
import {AsyncPipe} from '@angular/common';
import {TaskFormSmComponent} from '@client/features/project';

@Component({
    selector: "fr-board-view-feature",
    imports: [
        CdkDropList,
        CdkDrag,
        CdkDropListGroup,
        TaskCardComponent,
        ColumnHeaderComponent,
        AsyncPipe,
        CreateStatusFormComponent,
        TaskFormSmComponent,
    ],
    providers: [BoardViewApiService],
    styleUrl: './board-view.component.scss',
    template: `
        @if (columns$ | async; as columns) {
            <div cdkDropListGroup class="flex justify-start items-start no-wrap gap-4 w-full">
                @for (column of columns; track $index) {
                    <div class="card board-column">
                        <div class="flex justify-between mb-4">
                            <div class="text-slate-600 card-title">{{ column.title }}</div>
                            <fr-column-header [project]="project" [column]="column"/>
                        </div>

                        <div class="flex flex-col gap-2 min-h-2 my-4"
                             cdkDropList
                             [id]="column.id"
                             [cdkDropListData]="column.tasks"
                             (cdkDropListDropped)="drop($event)">
                            @for (task of column.tasks; track task.id) {
                                <fr-task-card cdkDrag [projectCode]="project.code"
                                              [task]="task"
                                              [cdkDragData]="task" />
                            }
                        </div>

                        <fr-task-form-sm [project]="project" [column]="column" />
                    </div>
                }
                <fr-create-status-form [projectId]="project.id" />
            </div>
        }
    `
})
export class BoardViewComponent implements OnInit {
    @Input() project: Project;
    private store = inject(Store<AppState>);

    columns$: Observable<StatusColumn[]>;

    ngOnInit(): void {
        const status$ = this.store.select(selectProjectStatusList(this.project.id));
        const tasks$ = this.store.select(selectTasksByProjectID(this.project.id));

        this.columns$ = combineLatest([status$, tasks$]).pipe(
            map(([statusList, tasks]) => {
                return statusList.map(status => ({
                    id: status.id.toString(),
                    title: status.title,
                    status: status,
                    tasks: tasks.filter(t => t.status_id === status.id)
                        .sort((a, b) => a.list_index - b.list_index)
                }));
            })
        );
    }

    drop(event: CdkDragDrop<Task[]>) {
        const container = event.container.data;
        const currentIdx = event.currentIndex;

        if (event.previousContainer === event.container) {
            moveItemInArray(container, event.previousIndex, currentIdx);
        } else {
            transferArrayItem(
                event.previousContainer.data,
                container,
                event.previousIndex,
                currentIdx
            );
        }

        const prev = container[currentIdx - 1];
        const next = container[currentIdx + 1];

        let newIndex: number;

        if (!prev && !next) {
            newIndex = 10000;
        } else if (!prev) {
            newIndex = next.list_index / 2;
        } else if (!next) {
            newIndex = prev.list_index + 10000;
        } else {
            newIndex = (prev.list_index + next.list_index) / 2;
        }

        this.store.dispatch(TaskChangeStatusAction({
            taskId: event.item.data.id,
            payload: {
                from_status: parseInt(event.previousContainer.id),
                to_status: parseInt(event.container.id),
                from_idx: event.previousIndex,
                to_idx: newIndex,
            }
        }));
    }
}
