import { Component } from "@angular/core";

@Component({
    selector: "app-auth-layout",
    imports: [],
    template: `
        <div
            style="height: 100vh; display: flex; align-items: center; justify-content: center"
        >
            <div>
                <img
                    src="img/logo-h.svg"
                    style="width: 10rem; margin-bottom: 1rem"
                />
                <div>
                    <ng-content />
                </div>
            </div>
        </div>
    `
})
export class AuthLayout {}
