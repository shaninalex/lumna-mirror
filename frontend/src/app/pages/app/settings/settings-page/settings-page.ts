import { Component, inject, OnInit } from '@angular/core';
import { UiService } from '@shared/ui';

@Component({
  selector: 'app-settings-page',
  imports: [],
  templateUrl: './settings-page.html',
})
export class SettingsPage implements OnInit {
    private ui = inject(UiService);

    ngOnInit() {
        this.ui.setPageTitle('Settings');
    }
}
