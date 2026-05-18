import { provideEffects } from "@ngrx/effects";
import { OnboardingApiService } from "./api";
import { OnboardingEffects } from "./model/onboarding.effects";

export function provideOnboardingFeature() {
    console.log("[PROVIDE] onboarding");
    return [provideEffects(OnboardingEffects), OnboardingApiService];
}
