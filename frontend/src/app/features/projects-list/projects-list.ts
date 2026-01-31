import { Component, inject, OnDestroy, OnInit, signal } from '@angular/core';
import { ProjectModel, ProjectPayload, ProjectState, selectProjects } from '@entities/project';
import { ProjectCard } from '@entities/project/ui';
import { AsyncPipe, NgClass } from '@angular/common';
import { ProjectForm } from '@features/project-form/project-form';
import { Dialog } from '@angular/cdk/dialog';
import { Store } from '@ngrx/store';
import { Observable, Subscription, tap } from 'rxjs';
import { actionProjectList } from '@entities/project/model/project.actions';

@Component({
    selector: 'app-projects-list-feature',
    imports: [ProjectCard, NgClass, AsyncPipe],
    templateUrl: './projects-list.html',
})
export class ProjectsListFeature implements OnInit, OnDestroy {
    private store = inject(Store<ProjectState>);
    projects: Observable<ProjectModel[]> = this.store.select(selectProjects);
    private s: Subscription = new Subscription();

    viewMode = signal<'list' | 'grid'>('grid');
    dialog = inject(Dialog);

    ngOnInit(): void {
        this.s.add(
            this.projects
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

    toggleViewMode(): void {
        if (this.viewMode() === 'list') {
            this.viewMode.set('grid');
        } else {
            this.viewMode.set('list');
        }
    }

    newProject(): void {
        const dialogRef = this.dialog.open<ProjectPayload>(ProjectForm, {
            width: '250px',
        });

        dialogRef.closed.subscribe((result) => {
            console.log(result);
            // if (result) this.dispatcher.dispatch(projectEvents.createProject(result));
        });
    }
}
