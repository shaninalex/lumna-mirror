import { Component, inject, OnDestroy, OnInit } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { ProjectState, selectProjects } from '@entities/project';
import { actionProjectList } from '@entities/project/model/project.actions';
import { Store } from '@ngrx/store';
import { Subscription, tap } from 'rxjs';

@Component({
    selector: 'app-projects-container',
    imports: [RouterOutlet],
    template: `
        <app-main-layout pageTitle="Projects">
            <router-outlet />
        </app-main-layout>
    `,
})
export class ProjectsContainer implements OnInit, OnDestroy {
    private store = inject(Store<ProjectState>);
    private s: Subscription = new Subscription();

    ngOnInit(): void {
        this.s.add(
            this.store
                .select(selectProjects)
                .pipe(
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
