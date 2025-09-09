import {Identity} from '@ory/kratos-client';


export interface UserModel {
    nickname: string
    identity: Identity
}
