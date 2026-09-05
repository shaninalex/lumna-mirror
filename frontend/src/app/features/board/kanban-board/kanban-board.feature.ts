import { AsyncPipe } from '@angular/common';
import type { OnInit } from '@angular/core';
import { Component, DestroyRef, effect, inject, input } from '@angular/core';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { Actions, ofType } from '@ngrx/effects';
import { Store } from '@ngrx/store';
import { actionsColumns, NewColumnFormComponent } from '@entities/column';
import { filter, tap, type Observable } from 'rxjs';
import { TimeAgoPipe } from '@shared/utils';
import { selectBoard, type BoardModel } from '@entities/board';
import { TaskCardComponent, actionTask, TaskInlineForm } from '@entities/task';
import type { CdkDragDrop } from '@angular/cdk/drag-drop';
import {
    CdkDrag,
    CdkDragHandle,
    CdkDropList,
    CdkDropListGroup,
} from '@angular/cdk/drag-drop';
import type { KanbanCard, KanbanColumn } from './model/kanban.models';
import { KanbanService } from './service';

@Component({
    selector: 'lu-kanban-board-feature',
    imports: [
        CdkDropListGroup,
        CdkDropList,
        CdkDrag,
        CdkDragHandle,
        AsyncPipe,
        NewColumnFormComponent,
        TimeAgoPipe,
        TaskCardComponent,
        TaskInlineForm,
    ],
    templateUrl: './kanban-board.feature.html',
    styleUrl: './kanban-board.feature.css',
    providers: [KanbanService],
})
export class KanbanBoardFeature implements OnInit {
    private store = inject(Store);
    private actions$ = inject(Actions);
    private destroyRef = inject(DestroyRef);
    private kanban = inject(KanbanService);

    boardId = input.required<number>();
    board$: Observable<BoardModel>;
    kolumns$: Observable<KanbanColumn[]> = this.kanban.boardData();

    constructor() {
        effect(() => {
            this.kanban.setBoardId(this.boardId());
        });
    }

    ngOnInit() {
        const _q = { board_id: this.boardId() };
        this.store.dispatch(actionTask.getList({ query: _q }));
        this.store.dispatch(actionsColumns.loadByBoardId(_q));
        this.board$ = this.store
            .select(selectBoard.byId(_q.board_id))
            .pipe(filter((board) => board !== null));

        this.actions$
            .pipe(
                ofType(actionsColumns.reorderFailed),
                takeUntilDestroyed(this.destroyRef),
                tap(() => this.store.dispatch(actionsColumns.loadByBoardId(_q))),
            )
            .subscribe();
    }

    dropColumn(event: CdkDragDrop<KanbanColumn[]>): void {
        this.kanban.dropColumn(event, this.boardId());
    }

    dropTask(event: CdkDragDrop<KanbanCard[]>, column: KanbanColumn): void {
        const isSameList = event.previousContainer === event.container;

        if (isSameList) {
            this.kanban.moveTask(event, column, this.boardId());
        } else {
            this.kanban.transferTask(event, column, this.boardId());
        }
    }

    public columnsAmount(): number {
        return this.kanban.getColumnsLength();
    }
}
