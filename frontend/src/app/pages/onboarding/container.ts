import { Component } from "@angular/core";
import { RouterOutlet } from "@angular/router";

@Component({
    selector: "onboarding-container",
    imports: [RouterOutlet],
    template: `
        <div
            class="vw-100 vh-100 d-flex align-items-center justify-content-center"
        >
            <div>
                <router-outlet />
                <div class="alert alert-info mt-3">
                    <p><b>make it via go templ!</b></p>
                    <p class="mb-0">Do NOT make temporary forms permanent</p>
                </div>
            </div>
        </div>
    `
})
export class OnboardingContainer {}
