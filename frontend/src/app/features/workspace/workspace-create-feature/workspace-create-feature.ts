import { Component, DestroyRef, inject, signal } from "@angular/core";
import { takeUntilDestroyed } from "@angular/core/rxjs-interop";

import { form, FormField, required } from "@angular/forms/signals";
import { Router } from "@angular/router";
import {
    actionWorkspaceCreate,
    actionWorkspaceCreateFailed,
    actionWorkspaceCreateSuccess,
    WorkspaceCreateModel
} from "@entities/workspace";
import { Actions, ofType } from "@ngrx/effects";
import { Store } from "@ngrx/store";


@Component({
    selector: "lu-workspace-create-feature",
    imports: [FormField],
    template: `
        <div class="container">
            <div>
                <form (submit)="onSubmit($event)">
                    <div class="mb-4">
                        <label for="workspace_title">Workspaces title</label>
                    </div>

                    <div class="mb-4">
                        <input type="text" class="form-control" id="workspace_title" [formField]="wspForm.title">
                    </div>

                    <div>
                        <button class="btn btn-primary" type="submit">Create</button>
                    </div>
                </form>
            </div>
        </div>
    `
})
export class WorkspaceCreateFeature {
    wspFormModel = signal<WorkspaceCreateModel>({ title: "" });
    wspForm = form(this.wspFormModel, (schemaPath) => {
        required(schemaPath.title);
    });
    private store = inject(Store);
    private actions$ = inject(Actions);
    private destroyRef = inject(DestroyRef);
    private router = inject(Router);

    constructor() {
        this.actions$
            .pipe(
                ofType(actionWorkspaceCreateSuccess),
                takeUntilDestroyed(this.destroyRef)
            )
            .subscribe(({ data }) => this.router.navigate(["/app", data.id]));

        this.actions$
            .pipe(
                ofType(actionWorkspaceCreateFailed),
                takeUntilDestroyed(this.destroyRef)
            )
            .subscribe((action) => console.log(action));
    }

    onSubmit(event: Event): void {
        event.preventDefault();
        this.store.dispatch(
            actionWorkspaceCreate({ data: this.wspFormModel() })
        );
    }
}
