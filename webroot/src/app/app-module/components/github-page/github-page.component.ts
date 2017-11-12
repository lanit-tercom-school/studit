import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Params } from '@angular/router';
import { Http, Headers, RequestOptions, Response } from '@angular/http';
import { environment } from '../../../../environments/environment';

@Component({
  selector: 'app-github-page',
  templateUrl: './github-page.component.html',
  styleUrls: ['./github-page.component.css']
})
export class GithubPageComponent implements OnInit {
  private code = '';

  constructor(private http: Http,
  private route: ActivatedRoute) { }

  ngOnInit() {
    this.route.queryParams
      .subscribe(params => {
        this.code = params['code'];
        console.log('get code ' + this.code);
        var query = `{ Token()
         }`;
         //*****************/
        this.http.get(environment.authUrl + '/graphql?query=' + query)
        .map((response: Response) =>  response.json().Token )
        .subscribe(res => {});

  });
  }
}
