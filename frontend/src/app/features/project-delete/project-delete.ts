import { DialogRef, DIALOG_DATA, Dialog } from '@angular/cdk/dialog';
import { AsyncPipe } from '@angular/common';
import { Component, inject, Input, OnInit } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { Router } from '@angular/router';
import {
    actionProjectDelete,
    actionProjectDeleteSuccess,
    ProjectModel,
    ProjectState,
} from '@entities/project';
import { Actions, ofType } from '@ngrx/effects';
import { Store } from '@ngrx/store';
import { filter, Observable, tap } from 'rxjs';

@Component({
    selector: 'app-project-delete-feature',
    imports: [AsyncPipe],
    template: `
        @if (project$ | async; as project) {
            <h2 class="mb-2">Danger</h2>
            <button class="btn btn-danger" (click)="openDialog(project.title)">Delete</button>
        }
    `,
})
export class ProjectDeleteFeature implements OnInit {
    private store = inject(Store<ProjectState>);
    private actions$ = inject(Actions);
    private router = inject(Router);
    private projectId: number;

    @Input() project$: Observable<ProjectModel>;

    dialog = inject(Dialog);

    ngOnInit() {
        this.project$.subscribe((project) => (this.projectId = project.id));
        this.actions$
            .pipe(ofType(actionProjectDeleteSuccess))
            .subscribe(() => this.router.navigate(['/projects']));
    }

    openDialog(title: string): void {
        const dialogRef = this.dialog.open<boolean>(DeleteProjectDialog, {
            data: title,
        });

        dialogRef.closed.subscribe((result) => {
            if (result) this.store.dispatch(actionProjectDelete({ project_id: this.projectId }));
        });
    }
}

@Component({
    selector: 'app-project-delete-feature-dialog',
    template: `
        <h1 class="text-lg font-bold">Are you sure want to delete project {{ projectName }}?</h1>
        <p class="mb-2">All data related to that project will be deleted</p>
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
    imports: [FormsModule],
    host: { class: 'modal' },
})
export class DeleteProjectDialog {
    dialogRef = inject<DialogRef<boolean>>(DialogRef<boolean>);
    projectName = inject(DIALOG_DATA);

    confirm(): void {
        this.dialogRef.close(true);
    }

    cancel(): void {
        this.dialogRef.close(false);
    }
}
