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
import { CardModule } from "primeng/card";
import { InputText } from "primeng/inputtext";
import { ButtonModule } from "primeng/button";

@Component({
    selector: "app-workspace-create-feature",
    imports: [FormField, CardModule, InputText, ButtonModule],
    template: `
        <div class="flex h-screen items-center justify-center">
            <p-card header="Create workspace">
                <form (submit)="onSubmit($event)">
                    <div class="mb-4">
                        <label for="workspace_title">Workspaces title</label>
                        <input
                            pInputText
                            id="workspace_title"
                            autocomplete="off"
                            class="block"
                            [class.p-invalid]="
                                wspForm.title().touched() &&
                                wspForm.title().invalid()
                            "
                            [formField]="wspForm.title"
                        />
                    </div>

                    <div>
                        <button pButton type="submit">Create</button>
                    </div>
                </form>
            </p-card>
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
