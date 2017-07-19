import { Injectable } from '@angular/core';
import { Http, Headers, RequestOptions, Response } from '@angular/http';
import { Observable } from 'rxjs/Observable';

import 'rxjs/add/operator/map';
import 'rxjs/add/operator/catch';
import 'rxjs/add/observable/throw';

import { environment } from '../../environments/environment';

import { User } from 'models/user';
import { UserRegister } from 'models/user-register';

@Injectable()
export class AuthService {

    constructor(private http: Http) { }

    authenticatenow(user: User) {
        var query: String = 'mutation{';
        query += 'Auth{ Signin(Login:"' + user.login;
        query += '" Password:"' + user.password + '")';
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
        var query = 'mutation{ Auth{';
        query += 'Activation(ActivationCode:"' + key + '")';
        query += '{ Message}}}';
        return this.http.get(environment.authUrl + '/graphql?query=' + query)
            .catch((error: any) => { return Observable.throw(error) });
    }

    register(user: UserRegister) {
        /*var headers = new Headers();
        headers.append('Content-Type', 'application/json');
        return this.http.post(environment.authUrl + '/v1/auth/signup/', JSON.stringify(user), { headers: headers })
            .map((res: Response) => {
                if (res.json().code)
                    localStorage.setItem('validation_code', res.json().code);
                else
                    return Observable.throw('no code');
            })
            .catch((error: any) => { return Observable.throw(error) });*/
        var query = 'mutation{Auth{Signup';
        query += '(Login:"' + user.login + '" ';
        query += 'Password: "' + user.password + '" ';
        query += 'Nickname: "' + user.nickname + '") ';
        query += '{ ActivationCode }}}';
        return this.http.get(environment.authUrl + '/graphql?query=' + query)
            .map((res: Response) => {
                var code = res.json().data.Auth.ActivationCode;
                if (code)
                    localStorage.setItem('validation_code', code);
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
