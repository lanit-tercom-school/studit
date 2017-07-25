
import { Injectable } from '@angular/core';
import { Http, Headers, Response } from '@angular/http';
import { Observable } from 'rxjs/Observable';
import 'rxjs/add/operator/map';
import 'rxjs/add/operator/catch';
import 'rxjs/add/observable/throw';

import { ProjectItem } from 'models/project-item';
import { UserInfo } from 'models/user-info';
import { environment } from '../../environments/environment';

@Injectable()
export class UserService {

  constructor(private http: Http) {
  }
  //TODO: It not work!
  getUsers() {
    return this.http.get(environment.apiUrl + '/v1/user/id/').map((response: Response) => response.json());
  }


  getUserById(id_: number): Observable<UserInfo> {
    var variable = { id: id_ };
    var query = `query($id:ID)  {
   User(Id:$id)
   {
    ID
    Nickname
    Avatar
    Description
  }
}&variables=`+ JSON.stringify(variable);
    return this.http.get(environment.apiUrl + '/graphql?query=' + query)
      .map((response: Response) => {
        return response.json().data.User;
      });

  }
//TODO: It not work!
  deleteUserById(id: number, token: string) {
    let headers = new Headers();
    headers.append('Accept', 'application/json');
    headers.append('Bearer-token', token);
    return this.http.delete(environment.apiUrl + '/v1/user/id/' + id, { headers: headers });
  }
  //TODO: It not work!
  changeUserById(id: number, token: string, user) {
    let headers = new Headers();
    headers.append('Content-Type', 'application/json');
    headers.append('Accept', 'application/json');
    headers.append('Bearer-token', token);
    return this.http.put(environment.apiUrl + '/v1/user/id/' + id, user, { headers: headers });
  }
  //TODO: It not work!
  changePasswordForUser(token: string, passwords) {
    let headers = new Headers();
    headers.append('Content-Type', 'application/json');
    headers.append('Accept', 'application/json');
    headers.append('Bearer-token', token);
    return this.http.put(environment.apiUrl + '/v1/auth/change-password/', passwords, { headers: headers });
  }

  // метод общий для студента и руководителя
  getProjectsOfUser(token: string, id_: number):Observable<ProjectItem[]> {
    var variable = { id: id_ };
    var query = `query($id:ID)  {
   User(Id:$id)
   {
     ProjectOn
     {
       Project
        {
          Id
          Description
          Name
          Logo
        }
      }
  }
}&variables=`+ JSON.stringify(variable);
    let headers = new Headers();
    headers.append('Authorization', 'Bearer ' + token);
    return this.http.get(environment.apiUrl + '/graphql?query=' + query, { headers: headers })
      .map((response: Response) => {
        let projects = new Array<ProjectItem>();
        response.json().data.User.ProjectOn.forEach(element => {
          projects.push(element.Project)
        });
        return projects
      });
  }
}
