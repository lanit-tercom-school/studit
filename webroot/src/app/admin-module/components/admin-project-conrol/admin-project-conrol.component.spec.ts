import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { AdminProjectConrolComponent } from './admin-project-conrol.component';

describe('AdminProjectConrolComponent', () => {
  let component: AdminProjectConrolComponent;
  let fixture: ComponentFixture<AdminProjectConrolComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ AdminProjectConrolComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(AdminProjectConrolComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should be created', () => {
    expect(component).toBeTruthy();
  });
});
