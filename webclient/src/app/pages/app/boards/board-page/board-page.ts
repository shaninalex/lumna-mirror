import {Component, inject, signal} from '@angular/core';
import {BoardModel} from '@entities/board';
import {ActivatedRoute, RouterLink} from '@angular/router';
import {CdkMenu, CdkMenuTrigger} from '@angular/cdk/menu';

@Component({
    selector: 'app-board-page',
    imports: [
        CdkMenu,
        CdkMenuTrigger,
        RouterLink
    ],
    template: `
        <div class="card bg-lime-200 flex items-top justify-between">
            <div>
                <div class="font-medium">{{ board()?.name }}</div>
            </div>

            <div>
                <button [cdkMenuTriggerFor]="menu">
                    <i class="fa-solid fa-ellipsis"></i>
                </button>

                <ng-template #menu>
                    <div class="bg-white border border-gray-200 rounded-xl p-4" cdkMenu>
                        <a [routerLink]="['edit']">Edit</a>
                    </div>
                </ng-template>
            </div>
        </div>
    `,
})
export class BoardPage {
    route = inject(ActivatedRoute);
    board = signal<BoardModel | null>(null);

    constructor() {
        this.route.data.subscribe(data => {
            this.board.set(data['board'])
        })
    }
}
