import { Component, inject } from '@angular/core';
import { Dialog, DialogRef } from '@angular/cdk/dialog';
import { Router, RouterLink } from "@angular/router";


@Component({
    selector: 'lu-switch-workspace',
    imports: [],
    template: `
        <button class="btn btn-sm btn-outline-secondary d-block w-100 text-left" (click)="openDialog()">
            <span class="d-block text-start">Lumna dev</span>
            <small class="d-block text-start" style="font-size: 0.8rem">Switch workspace</small>
        </button>
    `,
})
export class SwitchWorkspaceComponent {
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
            <a href="#" class="list-group-item list-group-item-action active" aria-current="true">
                The current link item
            </a>
            <a href="#" class="list-group-item list-group-item-action">A second link item</a>
            <a href="#" class="list-group-item list-group-item-action">A third link item</a>
            <a href="#" class="list-group-item list-group-item-action">A fourth link item</a>
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