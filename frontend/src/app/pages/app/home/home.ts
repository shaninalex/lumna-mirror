import { Component, inject, OnInit } from '@angular/core';
import { UiService } from '@shared/ui';

@Component({
    selector: 'app-home',
    templateUrl: './home.html',
    styleUrl: './home.css',
})
export class Home implements OnInit {
    private ui = inject(UiService);

    ngOnInit(): void {
        this.ui.setPageTitle('Overview');
    }
}
