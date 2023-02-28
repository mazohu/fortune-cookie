import { ComponentFixture, TestBed } from '@angular/core/testing';

import { EatCookieComponent } from './eat-cookie.component';

describe('EatCookieComponent', () => {
  let component: EatCookieComponent;
  let fixture: ComponentFixture<EatCookieComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ EatCookieComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(EatCookieComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
