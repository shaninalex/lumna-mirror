import { Component } from '@angular/core';
import { CdkMenu, CdkMenuItem, CdkMenuTrigger } from '@angular/cdk/menu';
import { SwitchWorkspaceComponent } from "../switch-workspace";


@Component({
    selector: 'lu-user-menu',
    imports: [CdkMenu, CdkMenuItem, CdkMenuTrigger, SwitchWorkspaceComponent],
    template: `
        <button [cdkMenuTriggerFor]="menu" class="btn btn-sm">
            <img src="images/7.png" alt="" class="rounded-circle" style="width: 24px">
        </button>
    
        <ng-template #menu>
            <div class="d-flex flex-column gap-2" cdkMenu>
                <div class="list-group">
                    <button cdkMenuItem type="button" class="list-group-item list-group-item-action">Account</button>
                    <button cdkMenuItem type="button" class="list-group-item list-group-item-action">Settings</button>
                    <button cdkMenuItem type="button" class="list-group-item list-group-item-action">Logout</button>
                </div>
                <lu-switch-workspace />
            </div>
        </ng-template>
    `,
})
export class UserMenuComponent {}
