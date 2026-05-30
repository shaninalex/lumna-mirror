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
    selector: "app-workspace-create-feature",
    imports: [FormField],
    template: `
        <div
            style="display:flex; height: 100vh; align-items: center; justify-content: center"
        >
            <form (submit)="onSubmit($event)">
                <div appearance="outlined">
                    <div>
                        <div>Poodle</div>
                        <div>Non-sporting group</div>
                    </div>
                    <div>
                        <div>
                            <div>Workspace title</div>
                            <input type="text" [formField]="wspForm.title" />
                            <div>
                                @if (
                                    wspForm.title().dirty() &&
                                    wspForm.title().errors()
                                ) {
                                    @for (
                                        error of wspForm.title().errors();
                                        track error
                                    ) {
                                        <div>{{ error.message }}</div>
                                    }
                                }
                            </div>
                        </div>
                    </div>
                    <div>
                        <button type="submit">Create</button>
                    </div>
                </div>
            </form>
        </div>
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
