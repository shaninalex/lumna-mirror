import { AsyncPipe } from '@angular/common';
import type { OnInit } from '@angular/core';
import { Component, inject, Input } from '@angular/core';
import type { ListModel } from '@entities/list';
import { selectListById } from '@entities/list/model/list.selectors';
import { Store } from '@ngrx/store';
import type { Observable } from 'rxjs';
import { filter } from 'rxjs';

@Component({
    selector: 'lu-board-detail-feature',
    imports: [AsyncPipe],
    templateUrl: './board-detail.feature.html',
})
export class BoardDetailFeature implements OnInit{
    private store = inject(Store);

    @Input() boardId: number;
    board$: Observable<ListModel>;

    ngOnInit() {
        this.board$ = this.store
            .select(selectListById(this.boardId))
            .pipe(filter((board) => board !== null));
    }
}
