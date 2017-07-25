/* tslint:disable:no-unused-variable */
import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { By } from '@angular/platform-browser';
import { DebugElement } from '@angular/core';

import { HomeStudentProjectsViewComponent } from './home-student-projects-view.component';

describe('HomeProjectsViewComponent', () => {
  let component: HomeStudentProjectsViewComponent;
  let fixture: ComponentFixture<HomeStudentProjectsViewComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ HomeStudentProjectsViewComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(HomeStudentProjectsViewComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});