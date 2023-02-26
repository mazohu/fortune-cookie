import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PastFortunesComponent } from './past-fortunes.component';

describe('PastFortunesComponent', () => {
  let component: PastFortunesComponent;
  let fixture: ComponentFixture<PastFortunesComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ PastFortunesComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(PastFortunesComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
