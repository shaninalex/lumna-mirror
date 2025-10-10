import {Component, inject, OnInit} from '@angular/core';
import {BoardViewComponent} from '@client/features/project';
import {ActivatedRoute} from '@angular/router';
import {Project} from '@client/entities/project';

@Component({
    selector: 'fr-board-view-page',
    imports: [
        BoardViewComponent,
    ],
    template: `
        <fr-board-view-feature [project]="project"/>
    `
})
export class BoardViewPageComponent implements OnInit {
    private route = inject(ActivatedRoute);
    project: Project

    ngOnInit() {
        this.route.data.subscribe(data => {
            this.project = data['project']
        })
    }
}
