import {Component, inject, Input, OnInit} from '@angular/core';
import {selectTasksByProjectID, Task, TaskCardComponent} from '@client/entities/task';
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
import {byMostRecent} from '@client/shared/common';

@Component({
    selector: "fr-board-view-feature",
    imports: [
        CdkDropList,
        CdkDrag,
        CdkDropListGroup,
        TaskCardComponent,
        ColumnHeaderComponent,
    ],
    providers: [BoardViewApiService],
    styleUrl: './board-view.component.scss',
    template: `
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
        </div>
    `
})
export class BoardViewComponent implements OnInit {
    @Input() project: Project;
    private boardApi = inject(BoardViewApiService);
    private store = inject(Store<AppState>);

    columns: StatusColumn[] = [];

    ngOnInit() {
        this.store.select(selectTasksByProjectID(this.project.id))
            .subscribe(tasks => {
                this.columns = this.project.statuses.map(status => ({
                    id: status.id,
                    title: status.title,
                    tasks: tasks.filter(t => t.status === status.id).sort(byMostRecent),
                }));
            });
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
