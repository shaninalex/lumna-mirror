import { patchState, signalStore, withMethods, withState } from '@ngrx/signals';
import { UserModel } from './user.model';

type UserState = {
    user: UserModel | null;
};

const initialState: UserState = {
    user: null,
};

export const UserStore = signalStore(
    { providedIn: 'root' },
    withState(initialState),
    withMethods((store) => ({
        setUser(user: UserModel | null): void {
            patchState(store, { user });
        },
    })),
);
