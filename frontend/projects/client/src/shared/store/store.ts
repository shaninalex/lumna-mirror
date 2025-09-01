import {SessionState} from '../../entities';
import {ProjectState} from '@client/entities/project/model/project.reducer';

export interface AppState {
    session: SessionState
    project: ProjectState
}
