import { Component, inject, OnInit } from '@angular/core';
import {ActivatedRoute, RouterLink} from '@angular/router';
import {BreadcrumbService} from '@shared/ui/breadcrumb.service';

@Component({
    selector: 'app-breadcrumbs-widget',
    template: `
        @if (breadcrumbs().length) {
            <nav aria-label="breadcrumb">
                <ol class="breadcrumb mb-0">
                    @for (crumb of breadcrumbs(); track crumb.url; let last = $last) {
                        <li
                            class="breadcrumb-item"
                            [class.active]="last"
                        >
                            @if (!last) {
                                <a [routerLink]="crumb.url">
                                    {{ crumb.label }}
                                </a>
                            } @else {
                                <span>{{ crumb.label }}</span>
                            }
                        </li>
                    }
                </ol>
            </nav>
        }
    `,
    imports: [
        RouterLink
    ]
})
export class BreadcrumbsWidget {
    private breadcrumbService = inject(BreadcrumbService);
    breadcrumbs = this.breadcrumbService.breadcrumbs;
}
