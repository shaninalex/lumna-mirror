import { Component, DestroyRef, inject, signal } from "@angular/core";
import { AppLayoutHeader } from "./app-layout-header";
import { AppLayoutSidebar } from "./app-layout-sidebar";
import { Store } from "@ngrx/store";
import { selectSidebarState } from "@core/store";
import { tap } from "rxjs";
import { takeUntilDestroyed } from "@angular/core/rxjs-interop";

@Component({
    selector: "app-layout",
    imports: [AppLayoutSidebar, AppLayoutHeader],
    templateUrl: "./app.layout.html",
    styleUrl: "./app.layout.css"
})
export class AppLayout {
    private store = inject(Store);
    private destroyRef = inject(DestroyRef);
    sidebarOpen = signal(true);

    private _ = this.store
        .select(selectSidebarState)
        .pipe(
            takeUntilDestroyed(this.destroyRef),
            tap((value) => this.sidebarOpen.set(value))
        )
        .subscribe();
}
