import { ComponentFixture, TestBed } from '@angular/core/testing';

import { NewColumnForm } from './new-column-form';

describe('NewColumnForm', () => {
  let component: NewColumnForm;
  let fixture: ComponentFixture<NewColumnForm>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [NewColumnForm]
    })
    .compileComponents();

    fixture = TestBed.createComponent(NewColumnForm);
    component = fixture.componentInstance;
    await fixture.whenStable();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
