import { AsyncPipe } from '@angular/common';
import type { OnInit } from '@angular/core';
import { Component, inject, input } from '@angular/core';
import { Store } from '@ngrx/store';
import type { StatusModel } from '@entities/status';
import { actionsStatuses, ColumnItemComponent, NewColumnFormComponent, selectStatuses } from '@entities/status';
import { filter, type Observable } from 'rxjs';
import { selectListById, type ListModel } from '@entities/list';
import { TimeAgoPipe } from '@shared/utils';

@Component({
    selector: 'lu-board-detail-feature',
    imports: [AsyncPipe, NewColumnFormComponent, TimeAgoPipe, ColumnItemComponent],
    templateUrl: './board-detail.feature.html',
})
export class BoardDetailFeature implements OnInit {
    private store = inject(Store);

    boardId = input.required<number>();
    board$: Observable<ListModel>;
    columns$: Observable<StatusModel[]>;

    ngOnInit() {
        this.store.dispatch(actionsStatuses.loadByListId({ list_id: this.boardId() }));
        this.board$ = this.store
            .select(selectListById(this.boardId()))
            .pipe(filter((board) => board !== null));
        this.columns$ = this.store.select(selectStatuses.byListId(this.boardId()));
    }
}
