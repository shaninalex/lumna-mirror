import {Component, inject, OnInit} from '@angular/core';
import {ActivatedRoute} from '@angular/router';
import {Project} from '@client/entities/project';
import {filter, map, take, tap} from 'rxjs';
import {ReactiveFormsModule} from '@angular/forms';
import {ProjectSettingsFeatureComponent} from '@client/features/project';

@Component({
    selector: 'lu-project-settings-page',
    imports: [
        ReactiveFormsModule,
        ProjectSettingsFeatureComponent,
    ],
    template: `
        <lu-project-settings-feature [project]="project"/>
    `
})
export class ProjectSettingsPageComponent implements OnInit {
    private route = inject(ActivatedRoute);
    project: Project

    ngOnInit() {
        this.route.data.pipe(
            take(1),
            map(data => data['project'] as Project),
            filter(project => !!project),
            tap(project => this.project = project),
        ).subscribe()
    }
}
