/* tslint:disable:no-unused-variable */

import { TestBed, async, inject } from '@angular/core/testing';
import { CoursecardService } from './coursecard.service';

describe('Service: Coursecard', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [CoursecardService]
    });
  });

  it('should ...', inject([CoursecardService], (service: CoursecardService) => {
    expect(service).toBeTruthy();
  }));
});
