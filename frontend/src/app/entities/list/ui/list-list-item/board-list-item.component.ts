import { Component, inject, Input } from '@angular/core';
import type { ListModel } from '../../model/list.model';
import { RouterLink } from '@angular/router';
import { AppRoutes } from '@core';

@Component({
    selector: 'lu-board-list-item',
    imports: [RouterLink],
    template: `
        <a [routerLink]="appRoutes.board(board.id)" class="text-decoration-none">
            <div class="d-flex justify-content-between align-items-start">
                <div class="me-3">
                    <h5 class="mb-1">{{ board.title }}</h5>

                    <p class="mb-2 text-muted">
                        Main board for feature development, bugs and improvements.
                    </p>

                    <div class="d-flex gap-3 small text-muted">
                        <span>42 cards</span>
                        <span>5 columns</span>
                        <span>Updated 2 hours ago</span>
                    </div>
                </div>

                <i class="fa-solid fa-chevron-right text-muted mt-1"></i>
            </div>
        </a>
    `,
    host: {
        class: 'list-group-item list-group-item-action py-3',
    },
})
export class BoardListItemComponent {
    @Input() board: ListModel;

    readonly appRoutes = inject(AppRoutes);
}
