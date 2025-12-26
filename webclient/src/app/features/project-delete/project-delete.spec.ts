import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ProjectDelete } from './project-delete';

describe('ProjectDelete', () => {
  let component: ProjectDelete;
  let fixture: ComponentFixture<ProjectDelete>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ProjectDelete]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ProjectDelete);
    component = fixture.componentInstance;
    await fixture.whenStable();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
