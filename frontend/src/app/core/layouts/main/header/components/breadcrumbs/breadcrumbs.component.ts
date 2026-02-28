import { Component, inject, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';

@Component({
    selector: 'app-breadcrumbs',
    template: `
        <nav class="breadcrumb has-bullet-separator is-small" aria-label="breadcrumbs">
            <ul class="ml-0">
                <li><a href="#">Lumna</a></li>
                <li><a href="#">Development</a></li>
                <li class="is-active"><a href="#">Create mock db cli command</a></li>
            </ul>
        </nav>
    `,
})
export class BreadCrumbs implements OnInit {
    private route = inject(ActivatedRoute);

    ngOnInit(): void {
    }
}
