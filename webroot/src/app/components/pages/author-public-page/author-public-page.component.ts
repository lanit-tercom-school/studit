import { Component, OnInit } from '@angular/core';
import { ApiService } from "../../../services/api.service";
import { ActivatedRoute } from "@angular/router";

@Component({
  selector: 'app-author-public-page',
  templateUrl: './author-public-page.component.html',
  styleUrls: ['./author-public-page.component.css']
})
export class AuthorPublicPageComponent implements OnInit {

  private authorData;

  constructor(private apiService: ApiService,
    private route: ActivatedRoute) { }


  ngOnInit() {
    this.route.params
      .subscribe(params => { this.authorData = this.apiService.getPublicAuthorInfoById(+params['id']); });
  }

}
