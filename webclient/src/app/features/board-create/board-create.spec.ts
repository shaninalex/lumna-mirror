import { ComponentFixture, TestBed } from '@angular/core/testing';

import { BoardCreate } from './board-create';

describe('BoardCreate', () => {
  let component: BoardCreate;
  let fixture: ComponentFixture<BoardCreate>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [BoardCreate]
    })
    .compileComponents();

    fixture = TestBed.createComponent(BoardCreate);
    component = fixture.componentInstance;
    await fixture.whenStable();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
