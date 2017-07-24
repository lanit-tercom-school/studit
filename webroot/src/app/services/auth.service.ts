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
        var variables = { login: user.login, password: user.password };
        var query = `mutation ($login: String $password: String)
    {
      Auth
      {
        Signin(Login: $login Password: $password)
         {
           DataOfExpiration
           PermissionLevel
           Token
           User
            { 
                Id
                Avatar
                Nickname
                Description
            }
         }
    }
    } &variables=`+ JSON.stringify(variables);
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

    validate(key_: string) {
        var variable = { key: key_ };
        var query = `mutation ($key: String )
    {
      Auth
      {
        Activation(ActivationCode: $key)
         {
          Message
         }
      }
    } &variables=`+ JSON.stringify(variable);
        return this.http.get(environment.authUrl + '/graphql?query=' + query)
            .catch((error: any) => { return Observable.throw(error) });
    }

    register(user: UserRegister) {
        var variables = { login: user.login, nickname: user.nickname, password: user.password };
        var query = `mutation ($login: String $nickname: String $password: String)
    {
      Auth
      {
        Signup(Login: $login Nickname: $nickname Password: $password)
         {
           ActivationCode
         }
    }
    } &variables=`+ JSON.stringify(variables);
        var query = 'mutation{Auth{Signup';
        query += '(Login:"' + user.login + '" ';
        query += 'Password: "' + user.password + '" ';
        query += 'Nickname: "' + user.nickname + '") ';
        query += '{ ActivationCode }}}';
        return this.http.get(environment.authUrl + '/graphql?query=' + query)
            .map((res: Response) => {
                var code = res.json().data.Auth.Signup.ActivationCode;
                if (code)
                    localStorage.setItem('validation_code', code);
                else
                    return Observable.throw('no code');
            })
            .catch((error: any) => { return Observable.throw(error) });
    }

}
