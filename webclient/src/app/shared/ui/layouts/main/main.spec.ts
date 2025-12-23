import {ComponentFixture, TestBed} from '@angular/core/testing';

import {MainLayout} from './mainLayout';
import {provideRouter} from '@angular/router';

describe('Main', () => {
    let component: MainLayout;
    let fixture: ComponentFixture<MainLayout>;

    beforeEach(async () => {
        await TestBed.configureTestingModule({
            imports: [MainLayout],
            providers: [
                provideRouter([]),
            ]
        })
            .compileComponents();

        fixture = TestBed.createComponent(MainLayout);
        component = fixture.componentInstance;
        await fixture.whenStable();
    });

    it('should create', () => {
        expect(component).toBeTruthy();
    });
});
