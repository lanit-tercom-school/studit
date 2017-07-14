  
import { Injectable } from '@angular/core';
import { Http, Headers,  Response } from '@angular/http';
import { Observable } from 'rxjs/Observable';
import 'rxjs/add/operator/map';
import 'rxjs/add/operator/catch';
import 'rxjs/add/observable/throw';

import { environment } from '../../environments/environment';

@Injectable()
export class UserService {

  constructor(private http: Http) {
  }
 
  getUsers() {
    return this.http.get(environment.apiUrl + '/v1/user/id/').map((response: Response) => response.json());
  }
  getUserById(id: number) {
    return this.http.get(environment.apiUrl + '/v1/user/id/' + id).map((response: Response) => response.json());
  }

  deleteUserById(id: number, token: string) {
    let headers = new Headers();
    headers.append('Accept', 'application/json');
    headers.append('Bearer-token', token);
    return this.http.delete(environment.apiUrl + '/v1/user/id/' + id, { headers: headers });
  }
  changeUserById(id: number, token: string, user) {
    let headers = new Headers();
    headers.append('Content-Type', 'application/json');
    headers.append('Accept', 'application/json');
    headers.append('Bearer-token', token);
    return this.http.put(environment.apiUrl + '/v1/user/id/' + id, user, { headers: headers });
  }
  changePasswordForUser(token: string, passwords) {
    let headers = new Headers();
    headers.append('Content-Type', 'application/json');
    headers.append('Accept', 'application/json');
    headers.append('Bearer-token', token);
    return this.http.put(environment.apiUrl + '/v1/auth/change-password/', passwords, { headers: headers });
  }
  // метод общий для студента и руководителя
  getProjectsOfUser(id: number) {
    return this.http.get(environment.apiUrl + '/v1/project/id/?user=' + id).map((response: Response) => response.json());
  }
}
