import { inject, Injectable, OnInit } from '@angular/core';
import { UiService } from '@client/shared/ui';

@Injectable()
export abstract class PageTitleSetter implements OnInit {
    protected ui: UiService = inject(UiService);

    abstract pageTitle: string;

    ngOnInit(): void {
        this.ui.setTitle(this.pageTitle);
    }

    public setTitle(text: string): void {
        this.ui.setTitle(text);
    }
}
