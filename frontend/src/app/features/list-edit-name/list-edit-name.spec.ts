import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ListEditName } from './list-edit-name';

describe('ListEditName', () => {
  let component: ListEditName;
  let fixture: ComponentFixture<ListEditName>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ListEditName]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ListEditName);
    component = fixture.componentInstance;
    await fixture.whenStable();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
