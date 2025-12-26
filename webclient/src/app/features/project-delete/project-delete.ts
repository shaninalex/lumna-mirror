import { DialogRef, DIALOG_DATA, Dialog } from '@angular/cdk/dialog';
import { Component, inject, Input } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { projectEvents } from '@entities/project';
import { Dispatcher } from '@ngrx/signals/events';

@Component({
    selector: 'app-project-delete-feature',
    imports: [],
    template: `
        <div class="border-b border-gray-200 mb-4"></div>
        <h2 class="mb-4">Danger</h2>
        <button class="bg-red-500 text-white rounded-lg px-4 py-2 cursor-pointer" (click)="openDialog()">Delete</button>

    `,
})
export class ProjectDeleteFeature {
    dialog = inject(Dialog);
    @Input() projectId: number

    private readonly dispatcher = inject(Dispatcher)

    openDialog(): void {
        const dialogRef = this.dialog.open<boolean>(CdkDialogOverviewExampleDialog);

        dialogRef.closed.subscribe(result => {
            if (result) {
                this.dispatcher.dispatch(projectEvents.deleteProject(this.projectId))
            }
        });
    }
}


@Component({
    selector: 'app-project-delete-feature-dialog',
    template: `
        <h1 class="text-lg font-bold">Are you sure?</h1>
        <p class="mb-2">All data related to that project will be deleted</p>
        <div class="flex gap-2">
            <button (click)="confirm()" class="bg-red-500 text-white rounded-lg px-2 py-1 cursor-pointer">delete</button>
            <button (click)="cancel()" class="bg-gray-400 text-white rounded-lg px-2 py-1 cursor-pointer">Cancel</button>
        </div>
    `,
    imports: [FormsModule],
    host: { 'class': 'modal' }
})
export class CdkDialogOverviewExampleDialog {
    dialogRef = inject<DialogRef<boolean>>(DialogRef<boolean>);
    data = inject(DIALOG_DATA);

    confirm(): void { this.dialogRef.close(true) }
    cancel(): void { this.dialogRef.close(false) }
}