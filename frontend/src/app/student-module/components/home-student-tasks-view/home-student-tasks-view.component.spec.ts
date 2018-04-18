import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { HomeStudentTasksViewComponent } from './home-student-tasks-view.component';

describe('HomeStudentTasksViewComponent', () => {
  let component: HomeStudentTasksViewComponent;
  let fixture: ComponentFixture<HomeStudentTasksViewComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ HomeStudentTasksViewComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(HomeStudentTasksViewComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
