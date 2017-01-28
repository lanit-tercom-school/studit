/* tslint:disable:no-unused-variable */

import { TestBed, async } from '@angular/core/testing';
import { ProjectCardListComponent } from './project-card-list.component';
import {ProgressService} from "../progress.service";

describe('Component: ProjectCardList', () => {
  it('should create an instance', () => {
    let component = new ProjectCardListComponent(ProgressService);
    expect(component).toBeTruthy();
  });
});
