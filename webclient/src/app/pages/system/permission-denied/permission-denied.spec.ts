import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PermissionDenied } from './permission-denied';

describe('PermissionDenied', () => {
  let component: PermissionDenied;
  let fixture: ComponentFixture<PermissionDenied>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [PermissionDenied]
    })
    .compileComponents();

    fixture = TestBed.createComponent(PermissionDenied);
    component = fixture.componentInstance;
    await fixture.whenStable();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
