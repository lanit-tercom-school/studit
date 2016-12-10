/* tslint:disable:no-unused-variable */

import { TestBed, async } from '@angular/core/testing';
import { NewsListComponent } from './news-list.component';
import {NewsService} from "../news.service";

describe('Component: NewsList', () => {
  it('should create an instance', () => {
    let component = new NewsListComponent(NewsService);
    expect(component).toBeTruthy();
  });
});
