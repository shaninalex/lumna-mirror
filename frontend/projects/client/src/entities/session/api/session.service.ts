import {Injectable} from '@angular/core';

@Injectable({
    providedIn: 'root'
})
export class SessionService {
    // Verify session method
    // Should send get request to the backend and return kratos session object except identity ( since it's will obtain
    // via user entity service ).
}
