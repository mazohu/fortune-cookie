import { ComponentFixture, TestBed } from '@angular/core/testing';

import { EatcookieComponent } from './eatcookie.component';

describe('EatcookieComponent', () => {
  let component: EatcookieComponent;
  let fixture: ComponentFixture<EatcookieComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ EatcookieComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(EatcookieComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
