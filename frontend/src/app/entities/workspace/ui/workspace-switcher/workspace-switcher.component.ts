import { Component, inject } from '@angular/core';
import { Dialog, DialogRef } from '@angular/cdk/dialog';
import { Router, RouterLink } from "@angular/router";
import { Store } from '@ngrx/store';
import { selectWorkspaceList } from '@entities/workspace';
import { AsyncPipe } from '@angular/common';


@Component({
    selector: 'lu-workspace-switcher',
    imports: [AsyncPipe, RouterLink],
    template: `
        @if (workspaces$ | async; as workspaces) {
            @if (workspaces.length) {
                <button class="btn btn-sm btn-outline-secondary d-block w-100 text-left" (click)="openDialog()">
                    <span class="d-block text-start">Lumna dev</span>
                    <small class="d-block text-start" style="font-size: 0.8rem">Switch workspace</small>
                </button>
            } @else {
                <a routerLink="/app/workspaces/create" class="btn btn-primary">Create workspace</a>
            }
        }
    `,
})
export class WorkspaceSwitcherComponent {
    private store = inject(Store);
    dialog = inject(Dialog);

    workspaces$ = this.store.select(selectWorkspaceList);

    openDialog() {
        this.dialog.open(SwitchWorkspaceModal, {
            minWidth: '300px',
        });
    }
}

@Component({
    selector: 'lu-switch-workspace-modal',
    imports: [AsyncPipe, RouterLink],
    template: `
        @if (workspaces$ | async; as workspaces) {
            @if (workspaces.length) {
                <div class="list-group">
                    @for (item of workspaces; track $index) {
                        <a href="#" class="list-group-item list-group-item-action">{{ item.title }}</a>
                    }
        
                    @if (workspaces.length > 5) {
                        <button (click)="seeAll()"
                            class="list-group-item list-group-item-action text-decoration-underline">
                            See all
                        </button>
                    }
                </div>
            } @else {
                <a routerLink="/workspaces/create" class="btn btn-primary">Create workspace</a>
            }
        }
    `,
})
export class SwitchWorkspaceModal {
    private dialogRef = inject(DialogRef);
    private router = inject(Router);
    private store = inject(Store);

    workspaces$ = this.store.select(selectWorkspaceList);

    async seeAll() {
        this.dialogRef.close();
        await this.router.navigate(['/workspaces']);
    }
}