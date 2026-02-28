import {inject, Injectable, signal} from '@angular/core';
import {ActivatedRoute, NavigationEnd, Router} from '@angular/router';
import {filter} from 'rxjs';


export interface Breadcrumb {
    label: string;
    url: string;
}

@Injectable({
    providedIn: 'root'
})
export class BreadcrumbService {
    private router = inject(Router);
    private route = inject(ActivatedRoute);
    private _breadcrumbs = signal<Breadcrumb[]>([]);

    public readonly breadcrumbs = this._breadcrumbs.asReadonly();

    constructor() {
        this.router.events.pipe(
            filter(event => event instanceof NavigationEnd)
        ).subscribe(() => {
            const root = this.route.root;
            const breadcrumbs = this.buildBreadcrumbs(root);
            this._breadcrumbs.set(breadcrumbs);
        })
    }

    private buildBreadcrumbs(route: ActivatedRoute, url = '', breadcrumbs: Breadcrumb[] = []): Breadcrumb[] {
        const children = route.children;

        if (children.length === 0) {
            return breadcrumbs;
        }

        for (const child of children) {
            const routeURL = child.snapshot.url.map(segment => segment.path).join('/');
            if (routeURL) {
                url += `/${routeURL}`;
            }

            const breadcrumbData = child.snapshot.data['breadcrumb'];

            if (breadcrumbData) {
                const label = typeof breadcrumbData === 'function' ? breadcrumbData(child.snapshot.data) : breadcrumbData;
                const last = breadcrumbs[breadcrumbs.length - 1];
                if (!last || last.label !== label) {
                    breadcrumbs.push({label, url});
                }
            }


            return this.buildBreadcrumbs(child, url, breadcrumbs);
        }

        return breadcrumbs;
    }
}
