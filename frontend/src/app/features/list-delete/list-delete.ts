import { Component, inject, input } from '@angular/core';
import { Dispatcher } from '@ngrx/signals/events';
import { listEvents } from '@entities/list/model/list.events';
import { Dialog, DIALOG_DATA, DialogRef } from '@angular/cdk/dialog';

@Component({
    selector: 'app-list-delete-feature',
    imports: [],
    template: `
        <button (click)="openDialog()" class="text-red-400 hover:underline">delete</button>
    `,
})
export class ListDeleteFeature {
    readonly listId = input.required<string>();
    readonly listName = input.required<string>();
    private readonly dispatcher = inject(Dispatcher);
    dialog = inject(Dialog);

    openDialog(): void {
        const dialogRef = this.dialog.open<boolean>(DeleteListDialog, { data: this.listName() });
        dialogRef.closed.subscribe((result) => {
            if (result) this.dispatcher.dispatch(listEvents.delete(this.listId()));
        });
    }
}

@Component({
    selector: 'app-list-delete-feature-dialog',
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
export class DeleteListDialog {
    dialogRef = inject<DialogRef<boolean>>(DialogRef<boolean>);
    listName = inject(DIALOG_DATA);

    confirm(): void {
        this.dialogRef.close(true);
    }
    cancel(): void {
        this.dialogRef.close(false);
    }
}
