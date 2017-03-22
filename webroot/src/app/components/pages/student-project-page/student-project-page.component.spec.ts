/* tslint:disable:no-unused-variable */
import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { By } from '@angular/platform-browser';
import { DebugElement } from '@angular/core';

import { SProjectPageComponent } from './student-project-page.component';

describe('SProjectPageComponent', () => {
  let component: SProjectPageComponent;
  let fixture: ComponentFixture<SProjectPageComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ SProjectPageComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(SProjectPageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
