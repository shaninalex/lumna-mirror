import { AsyncPipe } from '@angular/common';
import { Component, inject } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Store } from '@ngrx/store';
import { filter, map, type Observable } from 'rxjs';
import { BoardDetailFeature } from "@features";
import { MainLayout } from '@core/layout';

@Component({
    selector: 'lu-board-detail-page',
    imports: [MainLayout, AsyncPipe, BoardDetailFeature],
    template: `
        <lu-main-layout>
            @if (boardId$ | async; as boardId) {
                <div class="container-fluid py-4 h-100 d-flex flex-column">
                    <lu-board-detail-feature [boardId]="boardId" />
                </div>
            }
        </lu-main-layout>
    `,
})
export class BoardDetailPage {
    private store = inject(Store);
    private activeRoute = inject(ActivatedRoute);

    boardId$: Observable<number> = this.activeRoute.paramMap.pipe(
        map((params) => params.get('boardId')),
        filter((boardId) => boardId !== null),
        map((boardId) => Number(boardId)),
    );
}
