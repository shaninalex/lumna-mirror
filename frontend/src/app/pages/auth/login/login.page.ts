import type { OnInit } from '@angular/core';
import { Component, inject } from '@angular/core';
import { UiService } from '@shared/ui';
import { AuthLoginFeature } from "@features";

@Component({
    selector: 'lu-login-page',
    imports: [AuthLoginFeature],
    templateUrl: './login.page.html',
})
export class LoginPage implements OnInit {
    private ui = inject(UiService);

    ngOnInit(): void {
        this.ui.setPageTitle("Login")
    }
}
