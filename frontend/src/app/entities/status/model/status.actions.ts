import { createActionGroup, props } from '@ngrx/store';
import type { StatusPayloadModel, StatusModel } from './status.model';
import type { Error } from '@shared/models';


export const actionsStatuses = createActionGroup({
    source: "Status",
    events: {
        'load by list id': props<{ list_id: number }>(),
        'load by list id success': props<{ statuses: StatusModel[] }>(),
        'load by list id failed': props<{ errors: Error[] }>(),
        'create': props<{ payload: StatusPayloadModel }>(),
        'create success': props<{ status: StatusModel }>(),
        'create failed': props<{ errors: Error[] }>(),
    }
});

