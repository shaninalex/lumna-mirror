export enum InvitationModelState {
    pending,
    sent,
    accepted,
    validated,
    revoked
}

export interface InvitationModel {
    id: number;
    email: string;
    state: InvitationModelState;
    role: string;
    created_at: Date;
    valid_until: Date;
    meta: any;
}
