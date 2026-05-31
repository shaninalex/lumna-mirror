import { Component, inject } from "@angular/core";
import { RouterLink } from "@angular/router";
import { Store } from "@ngrx/store";
import { ProjectIcon, selectCurrentProjectList } from "@entities/project";
import { toSignal } from "@angular/core/rxjs-interop";
import { TagModule } from "primeng/tag";
import { ButtonModule } from "primeng/button";
import { InputTextModule } from "primeng/inputtext";
import { FormsModule } from "@angular/forms";
import { Checkbox } from "primeng/checkbox";

@Component({
    selector: "app-project-list",
    imports: [
        RouterLink,
        TagModule,
        ButtonModule,
        InputTextModule,
        FormsModule,
        Checkbox,
        ProjectIcon
    ],
    template: `<div class="p-4">
        <div class="rounded-md border border-zinc-200 bg-white p-5 mb-4">
            <p class="text-sm text-zinc-500">Total Projects</p>
            <div class="mt-3 flex items-end justify-between">
                <h2 class="text-3xl font-bold">{{ projects()?.length }}</h2>
                <p-tag
                    severity="success"
                    value="+2 this month"
                    [rounded]="true"
                />
            </div>
        </div>

        <!-- Projects -->
        <section
            class="overflow-hidden rounded-md border border-zinc-200 bg-white mb-4"
        >
            <div
                class="flex items-center justify-between border-b border-zinc-200 px-5 py-4"
            >
                <div>
                    <h2 class="text-lg font-semibold">Projects</h2>

                    <p class="text-sm text-zinc-500">
                        Recently active projects
                    </p>
                </div>

                <p-button label="Create" routerLink="create" />
            </div>

            <div
                class="flex items-center border-b border-zinc-200 px-5 py-4 gap-4"
            >
                <input
                    type="text"
                    placeholder="Search"
                    pInputText
                    [(ngModel)]="searchValue"
                />

                <div class="flex items-center">
                    <p-checkbox
                        inputId="archived"
                        name="pizza"
                        [binary]="true"
                        [(ngModel)]="archivedValue"
                    />
                    <label for="archived" class="ml-2"> Archived </label>
                </div>
            </div>
            <div class="divide-y divide-zinc-200">
                @for (project of projects(); track project.id) {
                    <a
                        [routerLink]="project.appLink"
                        class="flex items-center justify-between px-5 py-4 hover:bg-zinc-50"
                    >
                        <div class="flex items-center gap-2">
                            <app-project-icon [project]="project" />
                            <div>
                                <h3 class="font-medium">{{ project.title }}</h3>

                                <p class="text-sm text-zinc-500">
                                    42 open issues • 2 active boards
                                </p>
                            </div>
                        </div>

                        <p-tag
                            severity="success"
                            value="Active"
                            [rounded]="true"
                        />
                    </a>
                }
            </div>
        </section>
    </div>`
})
export class ProjectListPage {
    private store = inject(Store);
    searchValue = "";
    archivedValue = false;
    projects = toSignal(this.store.select(selectCurrentProjectList));
}
