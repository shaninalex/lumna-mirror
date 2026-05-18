export enum OnboardingState {
    WORKSPACES,
    COMPLETED
}

export interface UserOnboardingModel {
    email: string;
    first_name: string;
    last_name: string;
}

export enum InvitationProcessState {
    VALIDATING,
    SUCCESS,
    FAILED
}
