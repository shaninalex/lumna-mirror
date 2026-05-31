import { Component, DestroyRef, inject, OnInit } from "@angular/core";
import { selectWorkspaceList } from "@entities/workspace/model/workspace.selectors";
import { Store } from "@ngrx/store";
import { ButtonModule } from "primeng/button";
import { MenuModule } from "primeng/menu";
import { MenuItem } from "primeng/api";
import { tap } from "rxjs";
import { takeUntilDestroyed } from "@angular/core/rxjs-interop";

@Component({
    selector: "app-switch-workspaces",
    imports: [ButtonModule, MenuModule],
    template: `
        <p-menu [model]="items" #workspaceMenu [popup]="true" />
        <p-button
            (click)="workspaceMenu.toggle($event)"
            label="Switch workspaces"
            variant="outlined"
            size="small"
        />
    `
})
export class SwitchWorkspaces implements OnInit {
    private store = inject(Store);
    private destroyRef = inject(DestroyRef);
    items: MenuItem[] | undefined;

    ngOnInit() {
        this.store
            .select(selectWorkspaceList)
            .pipe(
                takeUntilDestroyed(this.destroyRef),
                tap((workspaces) => {
                    this.items = workspaces.map((workspace) => ({
                        label: workspace.title,
                        icon: "pi pi-file",
                        routerLink: `/app/${workspace.slug}`
                    }));
                })
            )
            .subscribe();
    }
}
