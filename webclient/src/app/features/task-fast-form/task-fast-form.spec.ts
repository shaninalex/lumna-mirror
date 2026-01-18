import { ComponentFixture, TestBed } from '@angular/core/testing';

import { TaskFastForm } from './task-fast-form';

describe('TaskFastForm', () => {
  let component: TaskFastForm;
  let fixture: ComponentFixture<TaskFastForm>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [TaskFastForm]
    })
    .compileComponents();

    fixture = TestBed.createComponent(TaskFastForm);
    component = fixture.componentInstance;
    await fixture.whenStable();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
