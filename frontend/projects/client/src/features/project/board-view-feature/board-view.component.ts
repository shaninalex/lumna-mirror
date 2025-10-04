import {Component, inject, Input, OnInit} from '@angular/core';
import {Task, TaskCardComponent} from '@client/entities/task';
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
import {map, Observable} from 'rxjs';
import {AsyncPipe} from '@angular/common';

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
    ],
    providers: [BoardViewApiService],
    styleUrl: './board-view.component.scss',
    template: `
        @if (columns$ | async; as columns) {
            <div cdkDropListGroup class="flex flex-row no-wrap gap-4">
                @for (column of columns; track column) {
                    <div class="bg-base-100 w-xs rounded-lg border border-base-300 p-4">
                        <fr-column-header [project]="project" [column]="column"/>
                        <div class="flex flex-col gap-2 min-h-20"
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
                    </div>
                }
                <div class="bg-base-100 w-xs rounded-lg border border-base-300 p-4">
                    Create status
                    <fr-create-status-form [projectId]="project.id" />
                </div>
            </div>
        }

    `
})
export class BoardViewComponent implements OnInit {
    @Input() project: Project;
    private boardApi = inject(BoardViewApiService);
    private store = inject(Store<AppState>);

    columns$: Observable<StatusColumn[]>;

    ngOnInit(): void {
        this.columns$ = this.store.select(selectProjectStatusList(this.project.id)).pipe(
            map(statusList => {
                return statusList.map(status => ({
                    id: status.id.toString(),
                    title: status.title,
                    status: status,
                    tasks: [],
                }))
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

        this.boardApi.ChangeStatus(this.project.code, event.item.data.code, {
            from_status: event.previousContainer.id,
            to_status: event.container.id,
            from_idx: event.previousIndex,
            to_idx: event.currentIndex,
        }).subscribe()
    }
}
