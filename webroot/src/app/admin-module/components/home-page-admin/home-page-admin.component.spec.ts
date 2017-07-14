import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { HomePageAdminComponent } from './home-page-admin.component';

describe('HomePageAdminComponent', () => {
  let component: HomePageAdminComponent;
  let fixture: ComponentFixture<HomePageAdminComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ HomePageAdminComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(HomePageAdminComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should be created', () => {
    expect(component).toBeTruthy();
  });
});
