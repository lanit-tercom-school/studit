import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { AdminUserConrolComponent } from './admin-user-conrol.component';

describe('AdminUserConrolComponent', () => {
  let component: AdminUserConrolComponent;
  let fixture: ComponentFixture<AdminUserConrolComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ AdminUserConrolComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(AdminUserConrolComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should be created', () => {
    expect(component).toBeTruthy();
  });
});
