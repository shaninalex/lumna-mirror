import {ComponentFixture, TestBed} from '@angular/core/testing';

import {Restore} from './restore';
import {provideRouter} from '@angular/router';

describe('Restore', () => {
    let component: Restore;
    let fixture: ComponentFixture<Restore>;

    beforeEach(async () => {
        await TestBed.configureTestingModule({
            imports: [Restore],
            providers: [
                provideRouter([])
            ]
        })
            .compileComponents();

        fixture = TestBed.createComponent(Restore);
        component = fixture.componentInstance;
        await fixture.whenStable();
    });

    it('should create', () => {
        expect(component).toBeTruthy();
    });
});
