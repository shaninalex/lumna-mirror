import { Component, inject, Input, OnInit } from '@angular/core';
import { BoardCreateFeature } from '@features/board-create';
import { RouterLink } from '@angular/router';
import { Store } from '@ngrx/store';
import { BoardModel, BoardState } from '@entities/board';
import { Observable } from 'rxjs';
import { selectBoardsByProjectId } from '@entities/board/model/board.selectors';
import { AsyncPipe } from '@angular/common';

@Component({
    selector: 'app-boards-list',
    imports: [BoardCreateFeature, RouterLink, AsyncPipe],
    template: `
        <div class="bg-(--color-secondary) card">
            <h2 class="font-medium text-lg mb-4">Boards</h2>
            @if (boards | async; as boards) {
                <nav class="flex flex-wrap gap-4">
                    @for (board of boards; track board.id) {
                        <a [routerLink]="['/projects', projectId, 'boards', board.id]" class="btn btn-primary">{{
                            board.title
                        }}</a>
                    }
                    <app-board-create-feature [projectId]="projectId" />
                </nav>
            }
        </div>
    `,
})
export class BoardsList implements OnInit {
    @Input() projectId: number;
    private store = inject(Store<BoardState>);
    boards: Observable<BoardModel[]>;

    ngOnInit() {
        this.boards = this.store.select(selectBoardsByProjectId(this.projectId));
    }
}
