import { AsyncPipe } from '@angular/common';
import type { OnInit } from '@angular/core';
import { Component, inject, Input } from '@angular/core';
import type { ListModel } from '@entities/list';
import { selectListById } from '@entities/list/model/list.selectors';
import { Store } from '@ngrx/store';
import type { Observable } from 'rxjs';
import { filter } from 'rxjs';
import { TaskCardComponent } from '@entities/task';
import { actionsStatuses, NewColumnFormComponent } from '@entities/status';

@Component({
    selector: 'lu-board-detail-feature',
    imports: [AsyncPipe, TaskCardComponent, NewColumnFormComponent],
    templateUrl: './board-detail.feature.html',
})
export class BoardDetailFeature implements OnInit {
    private store = inject(Store);

    @Input() boardId: number;
    board$: Observable<ListModel>;

    ngOnInit() {
        this.store.dispatch(actionsStatuses.loadByListId({ list_id: this.boardId }));
        this.board$ = this.store
            .select(selectListById(this.boardId))
            .pipe(filter((board) => board !== null));
    }
}
