import { Location } from '@angular/common';
import { inject, Injectable } from '@angular/core';
import { NavigationEnd, Router } from '@angular/router';
import { filter, tap } from 'rxjs';

@Injectable({
    providedIn: 'root',
})
export class NavigationService {
    private history: string[] = [];
    private router = inject(Router);
    private location = inject(Location);

    constructor() {
        this.router.events
            .pipe(
                filter((event) => event instanceof NavigationEnd),
                tap((event: NavigationEnd) => this.history.push(event.urlAfterRedirects)),
            )
            .subscribe();
    }

    back(): void {
        this.history.pop();

        if (this.history.length > 0) {
            this.location.back();
            return;
        }

        this.router.navigateByUrl('/');
    }
}
