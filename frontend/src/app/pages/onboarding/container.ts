import { Component } from "@angular/core";
import { RouterOutlet } from "@angular/router";

@Component({
    selector: "onboarding-container",
    imports: [RouterOutlet],
    template: `
        <div
            class="vw-100 vh-100 d-flex align-items-center justify-content-center"
        >
            <router-outlet />
        </div>
    `
})
export class OnboardingContainer {}
