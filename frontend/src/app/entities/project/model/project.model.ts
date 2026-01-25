import { BoardModel } from '@entities/board';

export interface ProjectModel {
    id: string;
    title: string;
    owner_id: string;
    boards: BoardModel[];
    created_at: Date;
    updated_at: Date;
}

// Used to create/patch projects
export interface ProjectPayload {
    title: string;
}
