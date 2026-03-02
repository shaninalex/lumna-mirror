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
import { Observable } from 'rxjs';

@Component({
    selector: 'app-project-delete-feature',
    imports: [AsyncPipe],
    template: `
        @if (project$ | async; as project) {
            <div class="alert alert-danger" role="alert">
                <h4 class="alert-heading">Danger!</h4>
                <p>This action will delete all project data</p>
                <hr>
                <button class="btn btn-danger" (click)="openDialog(project.title)">Delete</button>
            </div>
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
        <div class="card">
            <div class="card-body">
                <h5 class="fw-bold">Are you sure want to delete project {{ projectName }}?</h5>
                <p class="mb-2">All data related to that project will be deleted</p>
                <div class="d-flex gap-2">
                    <button (click)="confirm()" class="btn btn-danger">
                        Delete
                    </button>
                    <button (click)="cancel()" class="btn btn-secondary">
                        Cancel
                    </button>
                </div>
            </div>
        </div>

    `,
    imports: [FormsModule],
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
