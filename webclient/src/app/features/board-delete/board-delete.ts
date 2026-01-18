import {Component, computed, inject, Input, input} from '@angular/core';
import {BoardModel} from '@entities/board';
import {Dialog, DIALOG_DATA, DialogRef} from '@angular/cdk/dialog';
import {Dispatcher, Events} from '@ngrx/signals/events';
import {projectEvents, ProjectStore} from '@entities/project';
import {Router} from '@angular/router';
import {takeUntilDestroyed} from '@angular/core/rxjs-interop';
import {FormsModule} from '@angular/forms';
import {boardEvents} from '@entities/board/model/board.events';

@Component({
    selector: 'app-board-delete-feature',
    imports: [],
    template: `
        <h2 class="mb-4">Danger</h2>
        <button class="btn btn-danger" (click)="openDialog()">Delete</button>
    `,
})
export class BoardDeleteFeature {
    @Input() projectId: number
    dialog = inject(Dialog);
    board = input<BoardModel>();

    private readonly dispatcher = inject(Dispatcher)
    private readonly events = inject(Events)
    private router = inject(Router)

    constructor() {
        this.events
            .on(boardEvents._deleteSuccess)
            .pipe(takeUntilDestroyed())
            .subscribe(() => this.router.navigate(["/projects", this.projectId]));
    }

    openDialog(): void {
        const dialogRef = this.dialog.open<boolean>(DeleteBoardDialog, { data: this.board()?.name });
        dialogRef.closed.subscribe(result => {
            const b = this.board()
            if (result && b) this.dispatcher.dispatch(boardEvents.delete(b.id))
        });
    }
}


@Component({
    selector: 'app-board-delete-feature-dialog',
    template: `
        <h1 class="text-lg font-bold">Are you sure want to delete board {{ boardName }}?</h1>
        <p class="mb-2">All data (tasks) related to that board will be deleted</p>
        <div class="flex gap-2">
            <button (click)="confirm()" class="bg-red-500 text-white rounded-lg px-2 py-1 cursor-pointer">delete</button>
            <button (click)="cancel()" class="bg-gray-400 text-white rounded-lg px-2 py-1 cursor-pointer">Cancel</button>
        </div>
    `,
    imports: [FormsModule],
    host: { 'class': 'modal' }
})
export class DeleteBoardDialog {
    dialogRef = inject<DialogRef<boolean>>(DialogRef<boolean>);
    boardName = inject(DIALOG_DATA);

    confirm(): void { this.dialogRef.close(true) }
    cancel(): void { this.dialogRef.close(false) }
}
