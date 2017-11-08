import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { GithubPageComponent } from './github-page.component';

describe('GithubPageComponent', () => {
  let component: GithubPageComponent;
  let fixture: ComponentFixture<GithubPageComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ GithubPageComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(GithubPageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should be created', () => {
    expect(component).toBeTruthy();
  });
});
