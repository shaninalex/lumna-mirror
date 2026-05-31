import { Component, inject } from "@angular/core";
import { FormsModule } from "@angular/forms";
import { actionApplicationSidebarToggle } from "@core";
import { Store } from "@ngrx/store";
import { RippleModule } from "primeng/ripple";
import { AvatarModule } from "primeng/avatar";
import { BadgeModule } from "primeng/badge";
import { MenubarModule } from "primeng/menubar";
import { InputTextModule } from "primeng/inputtext";
import { ButtonModule } from "primeng/button";

@Component({
    selector: "app-app-layout-header",
    imports: [
        FormsModule,
        AvatarModule,
        BadgeModule,
        MenubarModule,
        InputTextModule,
        RippleModule,
        ButtonModule
    ],
    styleUrl: "./app-layout-header.css",
    template: `
        <div
            class="flex items-center py-2 px-4 gap-4 border-b border-b-slate-200"
        >
            <img src="/img/logo-h.svg" alt="" style="width: 150px" />
            <div class="grow"></div>
            <div class="flex gap-2">
                <p-button icon="pi pi-bell" outlined size="small" />
                <p-button icon="pi pi-cog" outlined size="small" />
                <div class="flex items-center gap-2">
                    <p-avatar image="/img/6.png" shape="circle" />
                </div>
            </div>
        </div>
    `
})
export class AppLayoutHeader {
    private store = inject(Store);
    // searchFormModel = signal<{ query: string }>({ query: "" });
    // searchForm = form(this.searchFormModel, (schemaPath) => {
    //     required(schemaPath.query, { message: "query field is required" });
    // });

    handleToggleSidebar(): void {
        this.store.dispatch(actionApplicationSidebarToggle());
    }
}
