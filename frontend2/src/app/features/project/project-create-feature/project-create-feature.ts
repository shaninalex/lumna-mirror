import { Component, DestroyRef, inject, OnInit, signal } from "@angular/core";
import { ButtonDirective } from "primeng/button";
import { Card } from "primeng/card";
import { InputText } from "primeng/inputtext";
import { form, FormField, required } from "@angular/forms/signals";
import {
    actionProjectCreate,
    actionProjectCreateFailed,
    actionProjectCreateSuccessful,
    ProjectCreateModel,
    selectProjectLink
} from "@entities/project";
import { Store } from "@ngrx/store";
import { Actions, ofType } from "@ngrx/effects";
import { Router } from "@angular/router";
import { selectWorkspaceCurrent } from "@entities/workspace";
import { takeUntilDestroyed } from "@angular/core/rxjs-interop";
import { tap } from "rxjs";
import { filter, switchMap } from "rxjs/operators";

@Component({
    selector: "app-project-create-feature",
    imports: [ButtonDirective, Card, InputText, FormField],
    template: `
        <p-card header="Create Project">
            <form (submit)="onSubmit($event)">
                <div class="mb-4">
                    <label for="workspace_title">Project title</label>
                    <input
                        pInputText
                        id="project_title"
                        autocomplete="off"
                        class="block"
                        [class.p-invalid]="
                            pForm.title().touched() && pForm.title().invalid()
                        "
                        [formField]="pForm.title"
                    />
                </div>

                <div>
                    <button pButton type="submit">Create</button>
                </div>
            </form>
        </p-card>
    `
})
export class ProjectCreateFeature implements OnInit {
    private store = inject(Store);
    private actions$ = inject(Actions);
    private destroyRef = inject(DestroyRef);
    private router = inject(Router);

    pFormModel = signal<ProjectCreateModel>({
        title: "",
        workspace_id: 0
    });
    pForm = form(this.pFormModel, (schemaPath) => {
        required(schemaPath.title);
    });

    ngOnInit() {
        this.actions$
            .pipe(
                ofType(actionProjectCreateSuccessful),
                takeUntilDestroyed(this.destroyRef),
                switchMap(({ project }) =>
                    this.store.select(selectProjectLink(project.id))
                ),
                tap((link) => this.router.navigate([link]))
            )
            .subscribe();

        this.actions$
            .pipe(
                ofType(actionProjectCreateFailed),
                takeUntilDestroyed(this.destroyRef)
            )
            .subscribe((action) => console.log(action));

        this.store
            .select(selectWorkspaceCurrent)
            .pipe(
                takeUntilDestroyed(this.destroyRef),
                filter((workspace) => !!workspace),
                tap((workspace) => {
                    this.pFormModel.set({
                        title: "",
                        workspace_id: workspace.id
                    });
                })
            )
            .subscribe();
    }

    onSubmit(event: Event): void {
        event.preventDefault();
        this.store.dispatch(actionProjectCreate({ data: this.pFormModel() }));
        this.formReset();
    }

    formReset(): void {
        this.pForm().reset();
        this.pFormModel.set({
            ...this.pFormModel(),
            title: ""
        });
    }
}
