import { Component, inject } from '@angular/core';
import { Dialog, DialogRef } from '@angular/cdk/dialog';
import { Router } from "@angular/router";


@Component({
    selector: 'lu-workspace-switcher',
    imports: [],
    template: `
        <button class="btn btn-sm btn-outline-secondary d-block w-100 text-left" (click)="openDialog()">
            <span class="d-block text-start">Lumna dev</span>
            <small class="d-block text-start" style="font-size: 0.8rem">Switch workspace</small>
        </button>
    `,
})
export class WorkspaceSwitcherComponent {
    dialog = inject(Dialog);

    openDialog() {
        this.dialog.open(SwitchWorkspaceModal, {
            minWidth: '300px',
        });
    }
}

@Component({
    selector: 'lu-switch-workspace-modal',
    template: `
        <div class="list-group">
            <a href="#" class="list-group-item list-group-item-action active">Lumna Dev</a>
            <a href="#" class="list-group-item list-group-item-action">Acme Inc.</a>
            <a href="#" class="list-group-item list-group-item-action">Open Source</a>
            <button (click)="seeAll()"
                class="list-group-item list-group-item-action text-decoration-underline">
                See all
            </button>
        </div>
    `,
})
export class SwitchWorkspaceModal {
    private dialogRef = inject(DialogRef);
    private router = inject(Router);

    async seeAll() {
        this.dialogRef.close();
        await this.router.navigate(['/workspaces']);
    }
}