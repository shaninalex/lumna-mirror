export enum OnboardingState {
    DISABLED,
    ALLOWED,
    COMPLETED
}

export interface OnboardingStateResponse {
    state: OnboardingState;
}

export interface UserOnboardingModel {
    email: string;
    first_name: string;
    last_name: string;
}

export enum InvitationProcessState {
    PENDING,
    SUCCESS,
    FAILED
}

export interface OnboardingContinue {
    email: string;
    state: string;
    role: string;
    meta: {
        first_name: string;
        last_name: string;
    };
}
