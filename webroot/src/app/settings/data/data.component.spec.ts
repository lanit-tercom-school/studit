/* tslint:disable:no-unused-variable */

import { TestBed, async } from '@angular/core/testing';
import { DataComponent } from './data.component';
import {SettingsService} from "../settings.service";

describe('Component: Data', () => {
  it('should create an instance', () => {
    let component = new DataComponent(SettingsService);
    expect(component).toBeTruthy();
  });
});
