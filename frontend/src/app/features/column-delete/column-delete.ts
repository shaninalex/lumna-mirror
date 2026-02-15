import { Component, inject, input } from '@angular/core';
import { Dialog, DIALOG_DATA, DialogRef } from '@angular/cdk/dialog';
import { Store } from '@ngrx/store';
import { actionColumnDelete, ColumnState } from '@entities/column';

@Component({
    selector: 'app-column-delete-feature',
    imports: [],
    template: `
        <button (click)="openDialog()" class="text-red-400 hover:underline cursor-pointer">
            delete
        </button>
    `,
})
export class ColumnDeleteFeature {
    private store = inject(Store<ColumnState>);
    columnId = input.required<number>();
    columnName = input.required<string>();
    dialog = inject(Dialog);

    openDialog(): void {
        const dialogRef = this.dialog.open<boolean>(DeleteColumnDialog, {
            data: this.columnName(),
        });
        dialogRef.closed.subscribe((result) => {
            if (result) this.store.dispatch(actionColumnDelete({ columnId: this.columnId() }));
        });
    }
}

@Component({
    selector: 'app-column-delete-feature-dialog',
    template: `
        <h1 class="text-lg font-bold">Are you sure want to delete "{{ listName }}" list?</h1>
        <p class="mb-2">All data (tasks) related to that list will be deleted too</p>
        <div class="flex gap-2">
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
    imports: [],
    host: { class: 'modal' },
})
export class DeleteColumnDialog {
    dialogRef = inject<DialogRef<boolean>>(DialogRef<boolean>);
    listName = inject(DIALOG_DATA);

    confirm(): void {
        this.dialogRef.close(true);
    }
    cancel(): void {
        this.dialogRef.close(false);
    }
}
