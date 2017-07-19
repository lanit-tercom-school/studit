import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { AdminNewsConrolComponent } from './admin-news-conrol.component';

describe('AdminNewsConrolComponent', () => {
  let component: AdminNewsConrolComponent;
  let fixture: ComponentFixture<AdminNewsConrolComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ AdminNewsConrolComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(AdminNewsConrolComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should be created', () => {
    expect(component).toBeTruthy();
  });
});
