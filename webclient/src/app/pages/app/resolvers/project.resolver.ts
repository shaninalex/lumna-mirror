import { ProjectModel, ProjectStore } from '@entities/project';
import { ActivatedRouteSnapshot, ResolveFn, RouterStateSnapshot } from '@angular/router';
import { inject } from '@angular/core';
import { toObservable } from '@angular/core/rxjs-interop';
import { filter, firstValueFrom, map, tap } from 'rxjs';
import { UiService } from '@shared/ui';

export const projectResolver: ResolveFn<ProjectModel> = (
    route: ActivatedRouteSnapshot,
    state: RouterStateSnapshot,
) => {
    const store = inject(ProjectStore);

    const ui = inject(UiService);
    const id = route.paramMap.get('id')!;

    let prefix = '';
    if (route.url.find((u) => u.path.includes('edit'))) {
        prefix = 'Edit ';
    }

    return firstValueFrom(
        toObservable(store.entities).pipe(
            filter((projects) => projects.some((p) => p.id === id)),
            map((projects) => projects.find((p) => p.id === id)!),
            tap((project) => {
                ui.setPageTitle(`${prefix}${project.title}`);
            }),
        ),
    );
};
