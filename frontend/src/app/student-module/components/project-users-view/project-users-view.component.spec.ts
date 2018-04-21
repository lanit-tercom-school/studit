import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ProjectUsersViewComponent } from './project-users-view.component';

describe('ProjectUsersViewComponent', () => {
  let component: ProjectUsersViewComponent;
  let fixture: ComponentFixture<ProjectUsersViewComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ProjectUsersViewComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ProjectUsersViewComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
