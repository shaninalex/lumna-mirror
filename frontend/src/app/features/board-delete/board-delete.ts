import { Component, inject, Input } from '@angular/core';
import { actionBoardDelete, actionBoardDeleteSuccess, BoardState } from '@entities/board';
import { Dialog, DIALOG_DATA, DialogRef } from '@angular/cdk/dialog';
import { Router } from '@angular/router';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { FormsModule } from '@angular/forms';
import { Actions, ofType } from '@ngrx/effects';
import { Store } from '@ngrx/store';
import { tap } from 'rxjs';

@Component({
    selector: 'app-board-delete-feature',
    imports: [],
    template: `
        <h2 class="mb-4">Danger</h2>
        <button class="btn btn-danger" (click)="openDialog()">Delete</button>
    `,
})
export class BoardDeleteFeature {
    @Input() boardId: number;
    @Input() boardTitle: string;

    dialog = inject(Dialog);

    private router = inject(Router);
    private actions$ = inject(Actions);
    private store = inject(Store<BoardState>);

    constructor() {
        // this.actions$
        //     .pipe(
        //         ofType(actionBoardDeleteSuccess),
        //         takeUntilDestroyed(),
        //         tap(() => this.router.navigate(['/projects', this.projectId])),
        //     )
        //     .subscribe();
    }

    openDialog(): void {
        const dialogRef = this.dialog.open<boolean>(DeleteBoardDialog, {
            data: this.boardTitle,
        });
        dialogRef.closed.subscribe((result) => {
            if (result) {
                this.store.dispatch(actionBoardDelete({ boardId: this.boardId }));
            }
        });
    }
}

@Component({
    selector: 'app-board-delete-feature-dialog',
    template: `
        <h1 class="text-lg font-bold">Are you sure want to delete board {{ boardName }}?</h1>
        <p class="mb-2">All data (tasks) related to that board will be deleted</p>
        <div class="d-flex gap-2">
            <button
                (click)="confirm()"
                class="bg-red-500 text-white rounded-lg px-2 py-1 cursor-pointer"
            >
                delete
            </button>
            <button
                (click)="cancel()"
                class="bg-gray-400 text-white rounded-lg px-2 py-1 cursor-pointer"
            >
                Cancel
            </button>
        </div>
    `,
    imports: [FormsModule],
    host: { class: 'modal' },
})
export class DeleteBoardDialog {
    dialogRef = inject<DialogRef<boolean>>(DialogRef<boolean>);
    boardName = inject(DIALOG_DATA);

    confirm(): void {
        this.dialogRef.close(true);
    }
    cancel(): void {
        this.dialogRef.close(false);
    }
}
