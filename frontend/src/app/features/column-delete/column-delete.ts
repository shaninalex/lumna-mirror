import { Component, inject, input } from '@angular/core';
import { Dialog, DIALOG_DATA, DialogRef } from '@angular/cdk/dialog';
import { Store } from '@ngrx/store';
import { actionColumnDelete, ColumnState } from '@entities/column';

@Component({
    selector: 'app-column-delete-feature',
    imports: [],
    template: `
        <button (click)="openDialog()" class="btn btn-link text-danger">
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
        <div class="card">
            <div class="card-body">
                <div class="fw-bold">Are you sure want to delete "{{ listName }}" list?</div>
                <p class="mb-2">All data (tasks) related to that list will be deleted too</p>
                <div class="d-flex gap-2">
                    <button (click)="confirm()" class="btn btn-danger">Delete</button>
                    <button (click)="cancel()" class="btn btn-secondary">Cancel</button>
                </div>
            </div>
        </div>
    `,
    imports: [],
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
