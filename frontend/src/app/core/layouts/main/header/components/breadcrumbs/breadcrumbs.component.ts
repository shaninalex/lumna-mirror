import { Component, inject, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';

@Component({
    selector: 'app-breadcrumbs',
    template: `
        <nav>
            <ol class="breadcrumb mb-0">
                <li class="breadcrumb-item"><a href="#">Home</a></li>
                <li class="breadcrumb-item"><a href="#">Library</a></li>
                <li class="breadcrumb-item active" aria-current="page">Data</li>
            </ol>
        </nav>
    `,
})
export class BreadCrumbs implements OnInit {
    private route = inject(ActivatedRoute);

    ngOnInit(): void {
    }
}
