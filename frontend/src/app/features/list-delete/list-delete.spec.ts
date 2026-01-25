import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ListDelete } from './list-delete';

describe('ListDelete', () => {
  let component: ListDelete;
  let fixture: ComponentFixture<ListDelete>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ListDelete]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ListDelete);
    component = fixture.componentInstance;
    await fixture.whenStable();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
