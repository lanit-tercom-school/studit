import { Injectable } from '@angular/core';
import { User } from 'models/user';
import { Http, Headers, RequestOptions, Response } from '@angular/http';
import { Observable } from 'rxjs/Observable';
import 'rxjs/add/operator/map';
import 'rxjs/add/operator/catch';
import 'rxjs/add/observable/throw';

import { environment } from '../../environments/environment';

@Injectable()
export class AuthService {

    constructor(private http: Http) { }

    authenticatenow(user: User) {
       /* return this.http.post(environment.authUrl + '/v1/auth/signin/', JSON.stringify(user))
            .map((response: Response) => {
                // successful login => getting jwt
                let res = response.json();
                if (res && res.bearer_token) {
                    // save data for keeping user logged in
                    res.login = user.login;
                    localStorage.setItem('current_user', JSON.stringify(res));
                }
            });*/
        var query: String = 'mutation{';
        query += 'Auth{ Signin(Login:"'+ user.login;
        query += '" Password:"'+ user.password + '")';
        query += '{ DataOfExpiration PermissionLevel Token ';
        query += 'User { Id Avatar Nickname Description }}}}';
        return this.http.get(environment.authUrl + '/graphql?query=' + query)
            .map((response: Response) => {
                // successful login => getting jwt
                let res = response.json().data.Auth.Signin;
                if (res && res.Token) {
                    // save data for keeping user logged in
                    res.User.Login = user.login;
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

    private jwt() {
        // create authorization-page header with jwt token
        let currentUser = JSON.parse(localStorage.getItem('current_user'));
        if (currentUser && currentUser.token) {
            let headers = new Headers({ 'authorization': 'Bearer ' + currentUser.token });
            return new RequestOptions({ headers: headers });
        }
    }
}
