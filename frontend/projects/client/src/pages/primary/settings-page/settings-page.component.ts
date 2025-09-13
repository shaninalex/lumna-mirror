import {Component} from '@angular/core';
import {SettingsFormComponent} from '@client/features/user';

@Component({
    selector: 'fr-settings-page',
    imports: [
        SettingsFormComponent
    ],
    template: `
        <fr-settings-form-feature />
    `,
})
export class SettingsPageComponent {

}
