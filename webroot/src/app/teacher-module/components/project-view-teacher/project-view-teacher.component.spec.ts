import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ProjectViewTeacherComponent } from './project-view-teacher.component';

describe('ProjectViewTeacherComponent', () => {
  let component: ProjectViewTeacherComponent;
  let fixture: ComponentFixture<ProjectViewTeacherComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ProjectViewTeacherComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ProjectViewTeacherComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should be created', () => {
    expect(component).toBeTruthy();
  });
});
