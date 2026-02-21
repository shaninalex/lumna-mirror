import { Component, inject, OnDestroy, OnInit } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { MainLayout } from '@core';
import { Store } from '@ngrx/store';
import { actionProjectList, ProjectState, selectProjects } from '@entities/project';
import { Subscription, take, tap } from 'rxjs';

@Component({
    selector: 'app-container',
    imports: [RouterOutlet, MainLayout],
    template: `
        <app-main-layout>
            <router-outlet />
        </app-main-layout>
    `,
})
export class DashboardContainer implements OnInit, OnDestroy {
    private store = inject(Store<ProjectState>);
    private s: Subscription = new Subscription();

    ngOnInit(): void {
        this.s.add(
            this.store
                .select(selectProjects)
                .pipe(
                    take(1),
                    tap((projects) => {
                        if (!projects.length) {
                            this.store.dispatch(actionProjectList());
                        }
                    }),
                )
                .subscribe(),
        );
    }

    ngOnDestroy(): void {
        this.s.unsubscribe();
    }
}
