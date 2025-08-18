import {Component} from '@angular/core';


@Component({
    selector: "ts-dropdown",
    template: `
    <!--
        <button matIconButton [cdkMenuTriggerFor]="project_detail_menu" class="example-standalone-trigger">
            <mat-icon>more_vert</mat-icon>
        </button>
        <ng-template #project_detail_menu>
            <div class="bg-white flex flex-col gap-2 border p-2 rounded" cdkMenu>
                <button cdkMenuItem>Refresh</button>
                <button cdkMenuItem>Settings</button>
                <button cdkMenuItem>Help</button>
                <button cdkMenuItem>Sign out</button>
            </div>
        </ng-template>

        It should used like this:

        <ts-dropdown>
            <ts-dropdown-button>
                <button>btn</button>
            </ts-dropdown-button>
            <ts-dropdown-content>
                <ul>
                    <li></li>
                    <li></li>
                    <li></li>
                    <li></li>
                </ul>
            </ts-dropdown-content>
        </ts-dropdown>
    -->
    `
})
export class UiDropdown {}
