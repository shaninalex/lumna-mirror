import { Component } from '@angular/core';
import { CdkMenu, CdkMenuItem, CdkMenuTrigger } from '@angular/cdk/menu';
import { RouterLink } from "@angular/router";


@Component({
    selector: 'lu-project-switcher',
    imports: [CdkMenu, CdkMenuItem, CdkMenuTrigger, RouterLink],
    template: `
        <button [cdkMenuTriggerFor]="menu" class="btn btn-sm btn-outline-secondary">
            Project name

            <i class="fa-solid fa-chevron-down"></i>
        </button>
    
        <ng-template #menu>
            <div class="list-group" cdkMenu>
                <button cdkMenuItem type="button" class="list-group-item d-flex justify-content-between align-items-center list-group-item-action gap-2">
                    <div class="me-auto">
                        Reports system
                    </div>
                    <span class="badge text-bg-primary rounded-pill">14</span>
                </button>
                <button cdkMenuItem type="button" class="list-group-item d-flex justify-content-between align-items-center list-group-item-action gap-2">
                    <div class="me-auto">
                        Inventory
                    </div>
                    <span class="badge text-bg-primary rounded-pill">14</span>
                </button>
                <button cdkMenuItem type="button" class="list-group-item d-flex justify-content-between align-items-center list-group-item-action gap-2">
                    <div class="me-auto">
                        SEO controll panel
                    </div>
                </button>
                <button cdkMenuItem type="button" class="list-group-item d-flex justify-content-between align-items-center list-group-item-action gap-2">
                    <div class="me-auto">
                        Ecommerce platform
                    </div>
                </button>
                
                <a cdkMenuItem routerLink="/projects" class="list-group-item list-group-item-action text-decoration-underline">
                    See all
                </a>
            </div>
        </ng-template>
    `,
})
export class ProjectSwitcherComponent { }
