import type { OnDestroy } from '@angular/core';
import { inject, Injectable, signal } from '@angular/core';
import { selectColumns } from '@entities/column';
import { selectTasks } from '@entities/task';
import { Store } from '@ngrx/store';
import type { Observable } from 'rxjs';
import { BehaviorSubject, combineLatest, filter, Subscription, switchMap, tap } from 'rxjs';
import type { KanbanColumn } from '../model/kanban.models';
import { toObservable } from '@angular/core/rxjs-interop';

@Injectable()
export class KanbanService implements OnDestroy {
    private store = inject(Store);
    private boardId = signal<number | undefined>(undefined);
    private sub = new Subscription();

    columns$ = toObservable(this.boardId).pipe(
        filter((id): id is number => id !== undefined),
        switchMap((id) => this.store.select(selectColumns.byListId(id))),
    );

    tasks$ = toObservable(this.boardId).pipe(
        filter((id): id is number => id !== undefined),
        switchMap((id) => this.store.select(selectTasks.byBoardId(id))),
    );

    private data: BehaviorSubject<KanbanColumn[]> = new BehaviorSubject<KanbanColumn[]>([]);

    ngOnDestroy(): void {
        this.sub.unsubscribe();
    }

    public setBoardId(boardId: number) {
        this.boardId.set(boardId);
        this.sub.add(
            combineLatest([this.columns$, this.tasks$])
                .pipe(
                    tap(([columns, tasks]) => {
                        const kolumns: KanbanColumn[] = []
                        columns.forEach((col) => {
                            const kolumn: KanbanColumn = { ...col, tasks: [] };
                            tasks.forEach((t) => {
                                const boardIndex = t.boards.findIndex(
                                    (b) => b.column_id === col.id && b.board_id === boardId,
                                );
                                if (boardIndex >= 0) {
                                    kolumn.tasks.push({
                                        ...t,
                                        column: col.id,
                                        position: t.boards[boardIndex].position,
                                    });
                                }
                            });
                            kolumns.push(kolumn);
                        });
                        this.data.next(kolumns);
                    }),
                )
                .subscribe(),
        );
    }

    public boardData(): Observable<KanbanColumn[]> {
        return this.data.asObservable();
    }

    public setData(d: KanbanColumn[]): void {
        return this.data.next(d);
    }

    public getData(): KanbanColumn[] {
        return this.data.getValue();
    }
}
