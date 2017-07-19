/* tslint:disable:no-unused-variable */

import { TestBed, async, inject } from '@angular/core/testing';
import { StudentService } from './student.service';

describe('AuthService', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [StudentService]
    });
  });

  it('should ...', inject([StudentService], (service: StudentService) => {
    expect(service).toBeTruthy();
  }));
});