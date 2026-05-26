import { Component, inject, signal } from "@angular/core";
import { form, required, FormField } from "@angular/forms/signals";
import {
    actionWorkspaceCreate,
    actionWorkspaceCreateFailed,
    actionWorkspaceCreateSuccess,
    WorkspaceCreateModel
} from "@entities/workspace";
import { Actions, ofType } from "@ngrx/effects";
import { Store } from "@ngrx/store";
import { tap } from "rxjs";

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
    private actions = inject(Actions);
    wspFormModel = signal<WorkspaceCreateModel>({ title: "" });
    wspForm = form(this.wspFormModel, (schemaPath) => {
        required(schemaPath.title);
    });

    constructor() {
        this.actions.pipe(
            ofType(actionWorkspaceCreateSuccess),
            tap((action) => console.log(action))
        );

        this.actions.pipe(
            ofType(actionWorkspaceCreateFailed),
            tap((action) => console.log(action))
        );
    }

    onSubmit(event: Event): void {
        event.preventDefault();
        this.store.dispatch(
            actionWorkspaceCreate({ data: this.wspFormModel() })
        );
    }
}
