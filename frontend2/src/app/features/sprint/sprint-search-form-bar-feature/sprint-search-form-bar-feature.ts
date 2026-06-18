import { Component } from "@angular/core";

@Component({
    selector: "app-sprint-search-form-bar-feature",
    imports: [],
    template: `
        <section
            class="mb-4 flex flex-col gap-2 rounded-md border border-zinc-200 bg-white p-4 md:flex-row md:items-center md:justify-between"
        >
            <div class="flex flex-1 items-center gap-3">
                <input
                    type="text"
                    placeholder="Search backlog..."
                    class="w-full rounded-md border border-zinc-300 bg-white px-4 py-2 text-sm outline-none focus:border-blue-500"
                />

                <button
                    class="rounded-md border border-zinc-300 px-4 py-2 text-sm hover:bg-zinc-50"
                >
                    Filters
                </button>
            </div>

            <div class="flex items-center gap-2 text-sm">
                <span class="rounded-full bg-zinc-200 px-3 py-1 text-zinc-700">
                    All Issues
                </span>

                <span class="rounded-full bg-blue-100 px-3 py-1 text-blue-700">
                    Sprint 12
                </span>
            </div>
        </section>
    `
})
export class SprintSearchFormBarFeature {}
