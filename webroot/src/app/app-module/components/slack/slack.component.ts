import { Component, OnInit } from '@angular/core';
import { Http, Headers, RequestOptions, Response } from '@angular/http';
import { ActivatedRoute } from '@angular/router';
import { Subscription } from 'rxjs/Subscription';

@Component({
  selector: 'app-slack',
  templateUrl: './slack.component.html',
  styleUrls: ['./slack.component.css']
})
export class SlackComponent implements OnInit {
  private client_id = '166136790673.166746603539';
  private secret = 'bb691d7438451edb2a0b3555f2170c6f';
  private code: string;
  constructor(private route: ActivatedRoute) {

    route.queryParams.subscribe(
      (queryParam: any) => {
        this.code = queryParam['code'];
        if(this.code){
          window.location.replace('https://slack.com/api/oauth.access?' + 'client_id=' + this.client_id + '&' + 'client_secret='+this.secret+ '&redirect_url=localhost:4200/slack&code='+this.code);
      }
      });
  }

  ngOnInit() {
  }
  auth() {
    window.location.replace('https://slack.com/oauth/authorize?' + 'client_id=' + this.client_id + '&' + 'scope=identity.basic&redirect_url=localhost:4200/slack');
  }
}
