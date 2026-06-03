import { Component, DestroyRef, inject, signal } from "@angular/core";
import { Store } from "@ngrx/store";
import { actionSprintCreate, actionSprintCreateSuccess, SprintCreateModel } from "@entities/sprint";
import { form, FormField, required } from "@angular/forms/signals";
import { Actions, ofType } from "@ngrx/effects";
import { InputTextModule } from "primeng/inputtext";
import { CardModule } from "primeng/card";
import { ButtonModule } from "primeng/button";
import { selectProjectCurrentId } from "@entities/project";
import { tap } from "rxjs";
import { takeUntilDestroyed } from "@angular/core/rxjs-interop";
import { filter } from "rxjs/operators";

@Component({
    selector: "app-task-inline-form-feature",
    imports: [],
    template: ` <p>task-inline-form-feature works!</p> `
})
export class TaskInlineFormFeature {
    private store = inject(Store);
    private actions$ = inject(Actions);
    private destroyRef = inject(DestroyRef);

    sprintFormModel = signal<SprintCreateModel>({
        title: "",
        project_id: 0
    });

    sprintForm = form(this.sprintFormModel, (schemaPath) => {
        required(schemaPath.title);
    });

    constructor() {
        this.store
            .select(selectProjectCurrentId)
            .pipe(
                takeUntilDestroyed(this.destroyRef),
                filter((projectId) => projectId !== undefined),
                tap((projectId) => {
                    this.sprintFormModel.set({
                        ...this.sprintFormModel(),
                        project_id: projectId
                    });
                })
            )
            .subscribe();
        this.actions$
            .pipe(
                ofType(actionSprintCreateSuccess),
                takeUntilDestroyed(this.destroyRef),
                tap(() => {
                    this.sprintFormModel.set({
                        ...this.sprintFormModel(),
                        title: ""
                    });
                    this.sprintForm().reset();
                })
            )
            .subscribe();
    }

    onSubmit(event: Event): void {
        event.preventDefault();
        // this.store.dispatch(
        //     actionTaskCreate({ data: this.sprintFormModel() })
        // );
    }
}

@Component({
    selector: "app-sprint-create-feature",
    imports: [InputTextModule, FormField, CardModule, ButtonModule],
    template: `
        <p-card>
            <form (submit)="onSubmit($event)" class="flex flex-col gap-2">
                <label for="sprintTitle">Sprint Title</label>
                <div class="flex gap-2">
                    <input
                        pInputText
                        id="sprintTitle"
                        class="grow"
                        [class.p-invalid]="
                            sprintForm.title().touched() &&
                            sprintForm.title().invalid()
                        "
                        [formField]="sprintForm.title"
                    />
                    <p-button label="Create" type="submit" />
                </div>

                <small> Enter sprint title </small>
            </form>
        </p-card>
    `
})
export class SprintCreateFeature {
    private store = inject(Store);
    private actions$ = inject(Actions);
    private destroyRef = inject(DestroyRef);

    sprintFormModel = signal<SprintCreateModel>({
        title: "",
        project_id: 0
    });

    sprintForm = form(this.sprintFormModel, (schemaPath) => {
        required(schemaPath.title);
    });

    constructor() {
        this.store
            .select(selectProjectCurrentId)
            .pipe(
                takeUntilDestroyed(this.destroyRef),
                filter((projectId) => projectId !== undefined),
                tap((projectId) => {
                    this.sprintFormModel.set({
                        ...this.sprintFormModel(),
                        project_id: projectId
                    });
                })
            )
            .subscribe();
        this.actions$
            .pipe(
                ofType(actionSprintCreateSuccess),
                takeUntilDestroyed(this.destroyRef),
                tap(() => {
                    this.sprintFormModel.set({
                        ...this.sprintFormModel(),
                        title: ""
                    });
                    this.sprintForm().reset();
                })
            )
            .subscribe();
    }

    onSubmit(event: Event): void {
        event.preventDefault();
        this.store.dispatch(
            actionSprintCreate({ data: this.sprintFormModel() })
        );
    }
}
