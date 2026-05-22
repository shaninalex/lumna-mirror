export interface UserModel {
    id: number;
    full_name: string;
    email: string;
    active: boolean;
    created_at: Date;
    updated_at: Date;
}

export interface AcceptInviteUserRegistrationForm {
    password: string;
    first_name: string;
    last_name: string;
    invitation_token: string;
}
