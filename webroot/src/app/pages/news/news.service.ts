import { Injectable } from '@angular/core';
import {NewsItem} from "./news-item";

@Injectable()
export class NewsService {

  constructor() { }

  getNews() : NewsItem[]{
    return [
      {
        "Id": 1,
        "Title": "Презентация студенческих проектов",
        "Body":"Презентация студенческих проектов.",
        "DateOfPublishing": new Date(),
        "Category":"Announcing"
      }
    ];
  }

}
