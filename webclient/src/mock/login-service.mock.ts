import { of, Subject } from 'rxjs';

export class LoginServiceMock {
    loginResult = of({});

    Login(_credentials: any) {
        return this.loginResult;
    }
}
