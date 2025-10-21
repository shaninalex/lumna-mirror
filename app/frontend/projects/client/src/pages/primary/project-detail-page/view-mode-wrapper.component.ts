import {Component, inject, OnInit} from '@angular/core';
import {ActivatedRoute, RouterLink, RouterOutlet} from '@angular/router';
import {Project} from '@client/entities/project';
import {BoardViewComponent} from '@client/features/project';

@Component({
    selector: `lu-project-overview-page`,
    imports: [
        RouterOutlet,
        BoardViewComponent,
        RouterLink
    ],
    template: `
        @switch (viewMode) {
            @case ('board') {
                <lu-board-view-feature [project]="project"/>
            }
            @default {
                <div class="text-red-600">View mode "{{ viewMode }}" is not ready</div>
                <a [routerLink]="['../']" [relativeTo]="route">Back</a>
            }
        }
        <router-outlet />
    `
})
export class ViewModeWrapperComponent implements OnInit {
    route = inject(ActivatedRoute)
    viewMode = this.route.snapshot.paramMap.get('viewMode');
    project: Project

    ngOnInit(): void {
        this.route.data.subscribe(data => {
            this.project = data['project']
        })
    }
}
