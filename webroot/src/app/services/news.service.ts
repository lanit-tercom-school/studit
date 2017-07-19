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

  // необязательные параметры
  getNewsPage(limit: number, offset: number) {
    if (limit > 0 && offset >= 0)
    {
      var query ='{NewsList(Offset:"' + offset + '" Limit: "' + limit + '")';
      query += '{Title Description DateOfCreation LastEdit Tags Image Id }}';
      return this.http.get(environment.apiUrl + '/graphql?query=' + query)
      .map((response: Response) => {return response.json().data});
   }
  }

  getNewsById(id_: number) {
    if (id_ >= 0 )
    {
      var query = '{News(Id: "' + id_ +'")';
      query +='{ Title Description DateOfCreation LastEdit Tags Image Id }}';
    return this.http.get(environment.apiUrl + '/graphql?query=' + query)
      .map((response: Response) => {return response.json().data.News});
    }
  }
}