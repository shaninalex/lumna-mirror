import { ComponentFixture, TestBed } from '@angular/core/testing';

import { BoardDelete } from './board-delete';

describe('BoardDelete', () => {
  let component: BoardDelete;
  let fixture: ComponentFixture<BoardDelete>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [BoardDelete]
    })
    .compileComponents();

    fixture = TestBed.createComponent(BoardDelete);
    component = fixture.componentInstance;
    await fixture.whenStable();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
