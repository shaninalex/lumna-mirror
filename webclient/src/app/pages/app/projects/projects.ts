import {Component} from '@angular/core';
import {MainLayout} from '@shared/ui/layouts/main/mainLayout';
import {ProjectsListFeature} from '@features/projects-list';

@Component({
    selector: 'app-projects',
    imports: [MainLayout, ProjectsListFeature],
    templateUrl: './projects.html',
    styleUrl: './projects.css',
})
export class Projects {

}
