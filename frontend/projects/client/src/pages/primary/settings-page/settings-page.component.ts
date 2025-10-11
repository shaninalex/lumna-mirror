import {Component} from '@angular/core';
import {SettingsFormComponent} from '@client/features/user';

@Component({
    selector: 'lu-settings-page',
    imports: [
        SettingsFormComponent
    ],
    template: `
        <lu-settings-form-feature />
    `,
})
export class SettingsPageComponent {

}
