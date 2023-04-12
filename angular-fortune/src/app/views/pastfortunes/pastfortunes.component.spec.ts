import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PastfortunesComponent } from './pastfortunes.component';

describe('PastfortunesComponent', () => {
  let component: PastfortunesComponent;
  let fixture: ComponentFixture<PastfortunesComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ PastfortunesComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(PastfortunesComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
