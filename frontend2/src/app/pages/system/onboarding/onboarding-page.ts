import { Component } from "@angular/core";
import { UserFormFeature } from "@features/onboarding";

@Component({
    selector: "app-onboarding-page",
    imports: [UserFormFeature],
    template: ` <div class="h-screen flex items-center justify-center">
        <div>
            <h1 class="mb-3 font-bold text-2xl">Onboarding</h1>
            <app-user-form-feature />
        </div>
    </div>`
})
export class OnboardingPage {}
