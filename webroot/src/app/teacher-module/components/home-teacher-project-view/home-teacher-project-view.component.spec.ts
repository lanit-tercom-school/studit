import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { HomeTeacherProjectViewComponent } from './home-teacher-project-view.component';

describe('HomeTeacherProjectViewComponent', () => {
  let component: HomeTeacherProjectViewComponent;
  let fixture: ComponentFixture<HomeTeacherProjectViewComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ HomeTeacherProjectViewComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(HomeTeacherProjectViewComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should be created', () => {
    expect(component).toBeTruthy();
  });
});
