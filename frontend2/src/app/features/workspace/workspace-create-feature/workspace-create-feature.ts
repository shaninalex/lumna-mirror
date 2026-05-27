import { Component, inject, DestroyRef, signal } from "@angular/core";
import { takeUntilDestroyed } from "@angular/core/rxjs-interop";

import { form, required, FormField } from "@angular/forms/signals";
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
    selector: "app-workspace-create-feature",
    imports: [FormField],
    template: `
        <form (submit)="onSubmit($event)">
            <div class="mb-3">
                <label for="title" class="text-sm">Workspace title</label>
                <input
                    class="form-control"
                    id="title"
                    type="text"
                    placeholder="Workspace title"
                    [formField]="wspForm.title"
                />
                @if (wspForm.title().dirty() && wspForm.title().errors()) {
                    @for (error of wspForm.title().errors(); track error) {
                        <div class="text-red-400 text-sm">
                            {{ error.message }}
                        </div>
                    }
                }
            </div>
            <button type="submit" class="btn">Login</button>
        </form>
    `
})
export class WorkspaceCreateFeature {
    private store = inject(Store);
    private actions$ = inject(Actions);
    private destroyRef = inject(DestroyRef);
    private router = inject(Router);

    wspFormModel = signal<WorkspaceCreateModel>({ title: "" });
    wspForm = form(this.wspFormModel, (schemaPath) => {
        required(schemaPath.title);
    });

    constructor() {
        this.actions$
            .pipe(
                ofType(actionWorkspaceCreateSuccess),
                takeUntilDestroyed(this.destroyRef)
            )
            .subscribe(({ data }) => this.router.navigate(["/app", data.slug]));

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
