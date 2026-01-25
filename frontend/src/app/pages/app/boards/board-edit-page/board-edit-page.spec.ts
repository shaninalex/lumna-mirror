import { ComponentFixture, TestBed } from '@angular/core/testing';

import { BoardEditPage } from './board-edit-page';

describe('BoardEditPage', () => {
  let component: BoardEditPage;
  let fixture: ComponentFixture<BoardEditPage>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [BoardEditPage]
    })
    .compileComponents();

    fixture = TestBed.createComponent(BoardEditPage);
    component = fixture.componentInstance;
    await fixture.whenStable();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
