import { Injectable } from '@angular/core';
import { Http, Headers, RequestOptions, Response } from '@angular/http';
import { Observable } from 'rxjs/Observable';
import 'rxjs/add/operator/map';
import 'rxjs/add/operator/catch';
import 'rxjs/add/observable/throw';

import { environment } from '../../environments/environment';

@Injectable()
export class NewsService {

  constructor(private http: Http) {
  }

  getNewsPage() {
    return this.http.get(environment.apiUrl + '/v1/news/').map((response: Response) => response.json());
  }

  getNewsById(id_: number) {
    return this.http.get(environment.apiUrl + '/v1/news/' + id_)
      .map((response: Response) => response.json());
  }
}