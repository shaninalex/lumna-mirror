export enum InvitationState {
    PENDING,
    SENT,
    ACCEPTED,
    REVOKED,
}

export interface InvitationModel {
    id: number;
    email: string;
    state: InvitationState;
    role: string;
    created_at: Date;
    valid_until: Date;
}
