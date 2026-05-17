import { Component } from "@angular/core";

@Component({
    selector: "app-app-layout-header",
    imports: [],
    template: `
        <header
            class="flex items-center justify-between border-b-[color:var(--color-slate-200)] h-[2.8rem] px-4 py-[0.6rem] border-b border-solid"
        >
            <div></div>
            <div class="relative">
                <input
                    class="form-control"
                    type="search"
                    placeholder="Search"
                />
                <i
                    class="fa-solid fa-magnifying-glass absolute right-2 top-2 text-slate-500"
                ></i>
            </div>
            <div class="flex items-center gap-3">
                <i class="fa-regular fa-bell"></i>
                <i class="fa-regular fa-circle-question"></i>
                <i class="fa-solid fa-gear"></i>
                <img src="img/6.png" class="w-8 h-8 rounded-full" alt="" />
            </div>
        </header>
    `,
    styles: `
        :host {
            position: absolute;
            top: 0;
            left: 0;
            width: 100%;
        }
    `
})
export class AppLayoutHeader {}
