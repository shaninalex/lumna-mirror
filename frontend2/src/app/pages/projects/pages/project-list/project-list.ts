import { Component } from "@angular/core";
import { RouterLink } from "@angular/router";

@Component({
    selector: "app-project-list",
    imports: [RouterLink],
    template: `<div class="p-4">
        <div class="rounded-md border border-zinc-200 bg-white p-5 mb-4">
            <p class="text-sm text-zinc-500">Total Projects</p>
            <div class="mt-3 flex items-end justify-between">
                <h2 class="text-3xl font-bold">12</h2>
                <span
                    class="rounded-full bg-green-100 px-2 py-1 text-xs font-medium text-green-700"
                >
                    +2 this month
                </span>
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

                <button class="font-bold text-blue-600 hover:text-blue-700">
                    Create
                </button>
            </div>

            <div class="divide-y divide-zinc-200">
                <!-- Project -->
                <a
                    routerLink="/app/lumna-1/project/lumna-new-frontend-13"
                    class="flex items-center justify-between px-5 py-4 hover:bg-zinc-50"
                >
                    <div class="flex items-center gap-2">
                        <div
                            class="flex h-12 w-12 items-center justify-center rounded-md bg-blue-600 text-sm font-bold text-white"
                        >
                            LM
                        </div>

                        <div>
                            <h3 class="font-medium">Lumna Core</h3>

                            <p class="text-sm text-zinc-500">
                                42 open issues • 2 active boards
                            </p>
                        </div>
                    </div>

                    <span
                        class="rounded-full bg-green-100 px-3 py-1 text-xs font-medium text-green-700"
                    >
                        Active
                    </span>
                </a>

                <!-- Project -->
                <a
                    routerLink="/app/lumna-1/project/sdondford-22"
                    class="flex items-center justify-between px-5 py-4 hover:bg-zinc-50"
                >
                    <div class="flex items-center gap-2">
                        <div
                            class="flex h-12 w-12 items-center justify-center rounded-md bg-purple-600 text-sm font-bold text-white"
                        >
                            API
                        </div>

                        <div>
                            <h3 class="font-medium">Public API</h3>

                            <p class="text-sm text-zinc-500">
                                17 open issues • 1 active sprint
                            </p>
                        </div>
                    </div>

                    <span
                        class="rounded-full bg-yellow-100 px-3 py-1 text-xs font-medium text-yellow-700"
                    >
                        Planning
                    </span>
                </a>

                <!-- Project -->
                <a
                    routerLink="/app/lumna-1/project/exelunatic-33"
                    class="flex items-center justify-between px-5 py-4 hover:bg-zinc-50"
                >
                    <div class="flex items-center gap-2">
                        <div
                            class="flex h-12 w-12 items-center justify-center rounded-md bg-pink-600 text-sm font-bold text-white"
                        >
                            UX
                        </div>

                        <div>
                            <h3 class="font-medium">UI Kit</h3>

                            <p class="text-sm text-zinc-500">
                                8 open issues • Design system
                            </p>
                        </div>
                    </div>

                    <span
                        class="rounded-full bg-blue-100 px-3 py-1 text-xs font-medium text-blue-700"
                    >
                        In Progress
                    </span>
                </a>
            </div>
        </section>
    </div>`
})
export class ProjectListPage {}
