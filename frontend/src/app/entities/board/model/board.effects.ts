import { inject, Injectable } from '@angular/core';
import { Actions } from '@ngrx/effects';

import { BoardApi } from '../api/board.api';

@Injectable()
export class BoardsEffects {
    private actions$ = inject(Actions);
    private projectsApi = inject(BoardApi);
}
