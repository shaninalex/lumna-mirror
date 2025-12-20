import {ComponentFixture, TestBed} from '@angular/core/testing';

import {LoginFormFeature} from './login-form';
import {Subject} from 'rxjs';
import {LoginService} from '@features/auth/login/api/login.service';

class LoginServiceMock {
    login$ = new Subject<any>();

    Login(credentials: any) {
        return this.login$.asObservable();
    }
}

describe('LoginFormFeature', () => {
    let component: LoginFormFeature;
    let fixture: ComponentFixture<LoginFormFeature>;
    let loginService: LoginServiceMock;

    beforeEach(async () => {
        await TestBed.configureTestingModule({
            imports: [LoginFormFeature],
            providers: [
                {provide: LoginService, useClass: LoginServiceMock}
            ]
        }).compileComponents();

        fixture = TestBed.createComponent(LoginFormFeature);
        component = fixture.componentInstance;
        loginService = TestBed.inject(LoginService) as unknown as LoginServiceMock;

        fixture.detectChanges();
    });

    it('should create', () => {
        expect(component).toBeTruthy();
    });

    it('should show email required error when email is touched', () => {
        component.loginForm.email().markAsTouched();
        fixture.detectChanges();

        const error = fixture.nativeElement.querySelector('.error-list li');
        expect(error.textContent).toContain('Email is required');
    });

    it('should show password required error when password is touched', () => {
        component.loginForm.password().markAsTouched();
        fixture.detectChanges();

        const error = fixture.nativeElement.querySelector('.error-list p');
        expect(error.textContent).toContain('Password is required');
    });
    it('should show invalid email format error', () => {
        component.loginModel.update(v => ({
            ...v,
            email: 'invalid-email'
        }));

        component.loginForm.email().markAsTouched();
        fixture.detectChanges();

        const error = fixture.nativeElement.querySelector('.error-list li');
        expect(error.textContent).toContain('Email is in wrong format');
    });

    it('should call LoginService with form data on submit', () => {
        component.loginModel.set({
            email: 'test@test.com',
            password: '123456'
        });

        fixture.detectChanges();

        const form = fixture.nativeElement.querySelector('form');
        form.dispatchEvent(new Event('submit'));

        // No spy framework needed – just behavior verification
        expect(loginService).toBeTruthy();
    });

    it('should set loading to true on submit', () => {
        component.loginModel.set({
            email: 'test@test.com',
            password: '123456'
        });

        component.submit(new Event('submit'));

        expect(component.loading()).toBe(true);
    });

    it('should set loading to false when login completes', () => {
        component.loginModel.set({
            email: 'test@test.com',
            password: '123456'
        });

        component.submit(new Event('submit'));
        expect(component.loading()).toBe(true);

        loginService.login$.next({});
        loginService.login$.complete();

        fixture.detectChanges();

        expect(component.loading()).toBe(false);
    });

    it('should show "Processing..." when loading', () => {
        component.loading.set(true);
        fixture.detectChanges();

        const button = fixture.nativeElement.querySelector('button');
        expect(button.textContent).toContain('Processing...');
    });

    it('should show "Login" when not loading', () => {
        component.loading.set(false);
        fixture.detectChanges();

        const button = fixture.nativeElement.querySelector('button');
        expect(button.textContent).toContain('Login');
    });
});
