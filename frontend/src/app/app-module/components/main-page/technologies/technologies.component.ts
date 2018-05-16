import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-technologies',
  templateUrl: './technologies.component.html',
  styleUrls: ['./technologies.component.css']
})
export class TechnologiesComponent implements OnInit {
  public Technologies = [
    {
      image: "./assets/technologies/angular.png",
      url: "https://angular.io/"
    },
    {
      image: "./assets/technologies/Golang.png",
      url: "https://golang.org/"
    },
    {
      image: "./assets/technologies/ts.png",
      url: "https://www.typescriptlang.org/"
    },
    {
      image: "./assets/technologies/postgreSQL.png",
      url: "https://www.postgresql.org/"
    },
    {
      image: "./assets/technologies/bootstrap.png",
      url: "https://getbootstrap.com/"
    },
    {
      image: "./assets/technologies/graphQL.png",
      url: "https://graphql.org/"
    },
    {
      image: "./assets/technologies/beego.png",
      url: "https://beego.me/"
    },
    {
      image: "./assets/technologies/swagger.png",
      url: "https://swagger.io/"
    },
    {
      image: "./assets/technologies/Jenkins.png",
      url: "https://jenkins.io/"
    },
    {
      image: "./assets/technologies/docker.png",
      url: "https://www.docker.com/"
    },
    {
      image: "./assets/technologies/Git.png",
      url: "https://git-scm.com/"
    },
    {
      image: "./assets/technologies/docker-compose.png",
      url: "https://docs.docker.com/compose/"
    }
  ];
  constructor() { }

  ngOnInit() {
  }

}
