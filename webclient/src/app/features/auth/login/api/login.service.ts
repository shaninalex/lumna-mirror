import {Injectable} from '@angular/core';
import {LoginCredentials} from '@features/auth/login/model/login.model';
import {Observable, of} from 'rxjs';

@Injectable({
    providedIn: 'root',
})
export class LoginService {
    public Login(payload: LoginCredentials): Observable<any> {
        /*
        * TODO: make real request
        * */

        return of(payload)
    }
}
