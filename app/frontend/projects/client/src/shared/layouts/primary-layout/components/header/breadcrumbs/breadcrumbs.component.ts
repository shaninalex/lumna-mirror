import {Component, inject, OnInit} from '@angular/core';
import {ActivatedRoute, NavigationEnd, Router, RouterLink} from '@angular/router';
import {filter} from 'rxjs';

interface Breadcrumb {
    label: string;
    url: string;
}

@Component({
    selector: "lu-breadcrumbs",
    imports: [
        RouterLink
    ],
    template: `
        <nav class="text-sm text-gray-600">
            @for (bc of breadcrumbs; track $index) {
                @if (!$last) {
                    <a [routerLink]="bc.url" class="hover:underline">{{ bc.label }}</a>
                    <span class="mx-1 text-gray-400">/</span>
                } @else {
                    <span aria-current="page">{{ bc.label }}</span>
                }
            }
        </nav>
    `
})
export class BreadcrumbsComponent implements OnInit {
    private router = inject(Router);
    private route = inject(ActivatedRoute);

    breadcrumbs: Breadcrumb[] = [];

    ngOnInit() {
        this.router.events
            .pipe(filter(e => e instanceof NavigationEnd))
            .subscribe(() => this.buildBreadcrumbs());
        this.buildBreadcrumbs();
    }

    private buildBreadcrumbs() {
        const breadcrumbs: Breadcrumb[] = [];
        let route: ActivatedRoute | null = this.route.root;
        let fullUrl = '';

        while (route) {
            const snapshot = route.snapshot;
            const config = snapshot.routeConfig;

            if (!config) {
                route = route.firstChild;
                continue;
            }

            const segment = snapshot.url.map(seg => seg.path).join('/') || config.path || '';
            if (segment) fullUrl += '/' + segment;

            const data = snapshot.data || {};
            const raw = config.data?.['breadcrumb'];
            let label: string | undefined;

            if (typeof raw === 'function') label = raw(data, snapshot.params, snapshot);
            else if (typeof raw === 'string') label = raw;
            else label =
                    (data as any)?.project?.title ??
                    snapshot.paramMap.keys.map(k => snapshot.paramMap.get(k)).find(Boolean) ??
                    segment;

            if (label) breadcrumbs.push({ label, url: fullUrl || '/' });
            route = route.firstChild;
        }

        this.breadcrumbs = breadcrumbs;
    }
}
