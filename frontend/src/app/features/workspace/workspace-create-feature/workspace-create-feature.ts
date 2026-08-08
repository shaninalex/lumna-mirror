import { Component, DestroyRef, inject, signal } from "@angular/core";
import { takeUntilDestroyed } from "@angular/core/rxjs-interop";
import { form, FormField, required } from "@angular/forms/signals";
import { actionWorkspaceCreate, actionWorkspaceCreateFailed, WorkspaceCreateModel } from "@entities/workspace";
import { Actions, ofType } from "@ngrx/effects";
import { Store } from "@ngrx/store";


@Component({
    selector: "lu-workspace-create-feature",
    imports: [FormField],
    template: `
        <div class="container">
            <form (submit)="onSubmit($event)">
                <div class="mb-4">
                    <label for="workspace_title" class="form-label">Workspaces title</label>
                    <input type="text" class="form-control" id="workspace_title" [formField]="wspForm.title">
                </div>

                <div>
                    <button class="btn btn-primary" type="submit">Create</button>
                </div>
            </form>
        </div>
    `
})
export class WorkspaceCreateFeature {
    private store = inject(Store);
    private actions$ = inject(Actions);
    private destroyRef = inject(DestroyRef);
    
    wspFormModel = signal<WorkspaceCreateModel>({ title: "" });
    wspForm = form(this.wspFormModel, (schemaPath) => required(schemaPath.title));
    
    constructor() {
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
