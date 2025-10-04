import {Component, inject, Input, OnInit} from '@angular/core';
import {ChangeTaskStatusAction, selectTasksByProjectID, Task, TaskCardComponent} from '@client/entities/task';
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
import {MatCardModule} from '@angular/material/card';
import {TaskFormSmComponent} from '@client/features/project/board-view-feature/components/task-form-sm';
import {byMostRecent} from '@client/shared/common';


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
        MatCardModule,
        TaskFormSmComponent,
    ],
    providers: [BoardViewApiService],
    styleUrl: './board-view.component.scss',
    template: `
        @if (columns$ | async; as columns) {
            <div cdkDropListGroup class="flex flex-row no-wrap gap-4">
                @for (column of columns; track column) {
                    <mat-card appearance="outlined" class="w-xs">
                        <mat-card-header class="flex justify-between items-start">
                            <mat-card-title class="text-slate-600">{{ column.title }}</mat-card-title>
                            <fr-column-header [project]="project" [column]="column"/>
                        </mat-card-header>
                        <mat-card-content>
                            <fr-task-form-sm [project]="project" [column]="column" />
                        </mat-card-content>
                        <mat-card-content>
                            <div class="flex flex-col gap-2 min-h-20 mt-4"
                                 cdkDropList
                                 [id]="column.id"
                                 [cdkDropListData]="column.tasks"
                                 (cdkDropListDropped)="drop($event)">
                                @for (task of column.tasks; track task.id) {
                                    <fr-task-card [projectKey]="project.code"
                                                  [task]="task"
                                                  [cdkDragData]="task"
                                                  cdkDrag/>
                                }
                            </div>
                        </mat-card-content>
                    </mat-card>
                }
                <mat-card appearance="outlined" class="w-xs">
                    <mat-card-header>
                        <mat-card-title class="text-slate-600">Create status</mat-card-title>
                    </mat-card-header>
                    <mat-card-content>
                        <fr-create-status-form [projectId]="project.id" />
                    </mat-card-content>
                </mat-card>
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
                        .sort(byMostRecent)
                        .sort((a, b) => (a.list_index - b.list_index))
                }));
            })
        );
    }

    drop(event: CdkDragDrop<Task[]>) {
        if (event.previousContainer === event.container) {
            moveItemInArray(event.container.data, event.previousIndex, event.currentIndex);
        } else {
            transferArrayItem(
                event.previousContainer.data,
                event.container.data,
                event.previousIndex,
                event.currentIndex,
            );
        }

        this.store.dispatch(ChangeTaskStatusAction({
            taskId: event.item.data.id,
            payload: {
                from_status: parseInt(event.previousContainer.id),
                to_status: parseInt(event.container.id),
                from_idx: event.previousIndex,
                to_idx: event.currentIndex,
            }
        }))
    }
}
