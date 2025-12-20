import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LoginFormFeature } from './login-form';

describe('LoginFormFeature', () => {
  let component: LoginFormFeature;
  let fixture: ComponentFixture<LoginFormFeature>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [LoginFormFeature]
    })
    .compileComponents();

    fixture = TestBed.createComponent(LoginFormFeature);
    component = fixture.componentInstance;
    await fixture.whenStable();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
