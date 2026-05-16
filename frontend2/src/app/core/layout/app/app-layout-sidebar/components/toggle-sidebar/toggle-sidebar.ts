import { AsyncPipe } from "@angular/common";
import { Component, inject } from "@angular/core";
import {
    actionApplicationSidebarToggle,
    selectSidebarState
} from "@core/store/app";
import { Store } from "@ngrx/store";

@Component({
    selector: "app-toggle-sidebar",
    imports: [AsyncPipe],
    template: `
        <button
            class="inline-block border border-slate-300 px-1 rounded hover:bg-slate-200 cursor-pointer"
            (click)="onClick()"
        >
            @if (state$ | async; as st) {
                <i class="fa-solid fa-chevron-right"></i>
            } @else {
                <i class="fa-solid fa-chevron-left"></i>
            }
        </button>
    `
})
export class ToggleSidebar {
    private store = inject(Store);
    state$ = this.store.select(selectSidebarState);

    onClick(): void {
        this.store.dispatch(actionApplicationSidebarToggle());
    }
}
