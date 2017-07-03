import { Injectable } from '@angular/core';
import { User } from 'models/user';
import { Http, Headers, Response } from '@angular/http';
import { Observable } from 'rxjs/Observable';
import 'rxjs/add/operator/map';
import 'rxjs/add/operator/catch';
import 'rxjs/add/observable/throw';

import { environment } from '../../environments/environment';

@Injectable()
export class AuthService {

    constructor(private http: Http) { }

    authenticatenow(user: User) {
        return this.http.post(environment.authUrl + '/v1/auth/signin/', JSON.stringify(user))
            .map((response: Response) => {
                // successful login => getting jwt
                let res = response.json();
                if (res && res.bearer_token) {
                    // save data for keeping user logged in
                    res.login = user.login;
                    localStorage.setItem('current_user', JSON.stringify(res));
                }
            });
    }
    unauthentificatenow() {
        localStorage.removeItem("current_user");
    }
      validate(key: string) {
    return this.http.get(environment.authUrl + '/v1/auth/signup/?pass=' + key)
      .catch((error: any) => { return Observable.throw(error) });
  }

  register(user) {
    var headers = new Headers();

    headers.append('Content-Type', 'application/json');

    return this.http.post(environment.authUrl + '/v1/auth/signup/', JSON.stringify(user), { headers: headers })
      .map((res: Response) => {
        if (res.json().code)
          localStorage.setItem('validation_code', res.json().code);
        else
          return Observable.throw('no code');
      })
      .catch((error: any) => { return Observable.throw(error) });
  }
}
