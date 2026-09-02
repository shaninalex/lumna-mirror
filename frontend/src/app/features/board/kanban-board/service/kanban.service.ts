import type { OnDestroy} from '@angular/core';
import { inject, Injectable, signal } from '@angular/core';
import { selectColumns } from '@entities/column';
import { selectTasks } from '@entities/task';
import { Store } from '@ngrx/store';
import type { Observable} from 'rxjs';
import { BehaviorSubject, combineLatest, filter, Subscription, switchMap, tap } from 'rxjs';
import type { KanbanCard, KanbanColumn } from '../model/kanban.models';
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

    private data: BehaviorSubject<KanbanColumn[]> = new BehaviorSubject<KanbanColumn[]>([])

    constructor() {
        console.log('Create KanbanService');
    }

    ngOnDestroy(): void {
        this.sub.unsubscribe();
    }

    public setBoardId(boardId: number) {
        this.boardId.set(boardId);
        this.sub.add(
            combineLatest([this.columns$, this.tasks$])
                .pipe(
                    tap(([columns, tasks]) => {
                        columns.forEach((col) => {
                            const kolumn: KanbanColumn = { ...col, tasks: [] };
                            // TODO: There are something incorrect here.
                            tasks.forEach((t) => {
                                const boardIndex = t.boards.findIndex((b) => b.column_id === col.id && b.board_id === col.board_id);
                                if (boardIndex > 0) {
                                    const kard: KanbanCard = {
                                        id: t.id,
                                        column: col.id,
                                        position: t.boards[boardIndex].position,
                                        task: t,
                                    };
                                    kolumn.tasks.push(kard);
                                }
                            });
                            const v = this.data.value;
                            v.push(kolumn)
                            this.data.next(v)
                        });
                    }),
                )
                .subscribe(),
        );
    }

    public boardData(): Observable<KanbanColumn[]> {
        return this.data.asObservable();
    }
}
