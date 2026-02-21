import { Component, inject, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';

@Component({
    selector: 'app-breadcrumbs',
    template: ``,
})
export class BreadCrumbs implements OnInit {
    private route = inject(ActivatedRoute);

    ngOnInit(): void {
        console.log('breadcrumbs');
    }
}
