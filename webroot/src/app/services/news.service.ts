import { Injectable } from '@angular/core';
import { Http, Headers, RequestOptions, Response } from '@angular/http';
import { Observable } from 'rxjs/Observable';

import 'rxjs/add/operator/map';
import 'rxjs/add/operator/catch';
import 'rxjs/add/observable/throw';

import { NewsList } from "models/news-list";
import { NewsItem } from "models/news-item";
import { environment } from '../../environments/environment';

@Injectable()
export class NewsService {

  constructor(private http: Http) {
  }

  // необязательные параметры
  getNewsPage(limit_: number, offset_: number):Observable<NewsList> {
    if (limit_ > 0 && offset_ >= 0) {
      var variables = { limit: limit_, offset: offset_ }
      var query = `query($limit:String, $offset: String)
      {
      NewsList(Offset: $offset Limit: $limit)
      {
        TotalCount
        NewsList
        {
          Title
          Description
          Created
          LastEdit
          Tags
          Image
          Id
        }
    }
  }&variables=`+ JSON.stringify(variables);
      return this.http.get(environment.apiUrl + '/graphql?query=' + query)
        .map((response: Response) =>  response.json().data.NewsList)
        .catch((error: any) => {
          return Observable.throw(error);
        });
    }
  }

  getNewsById(id_: number):Observable<NewsItem> {
    if (id_ >= 0 )
    {
      var variable = {id: id_};
      var query = `query($id: String)
      {
        News(Id: $id)
      {
        Title
        Description
        Created
        LastEdit
        Tags
        Image
        Id
      }
    }&variables=`+ JSON.stringify(variable);
    return this.http.get(environment.apiUrl + '/graphql?query=' + query)
      .map((response: Response) => {return response.json().data.News})
      .catch((error: any) => {
          return Observable.throw(error);
        });
    }
  }
}