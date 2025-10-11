import {BehaviorSubject, Observable} from 'rxjs';
import {ToastMessage} from '@client/shared/models';
import { v4 as uuidV4} from 'uuid'

export class MessagesClass {
    private _messages$: BehaviorSubject<ToastMessage[]> = new BehaviorSubject([] as ToastMessage[]);

    public messages$(): Observable<ToastMessage[]> {
        return this._messages$.asObservable()
    }

    public addMessage(msg: string): void {
        const messages = this._messages$.value
        messages.push({
            id: uuidV4(),
            message: msg,
        })
        this._messages$.next(messages)
    }

    public removeMessage(id: string): void {
        let messages = [...this._messages$.value]
        this._messages$.next(messages.filter(m => m.id !== id))
    }
}
