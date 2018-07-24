import { Injectable } from '@angular/core';
import { Http, Headers, RequestOptions, Response } from '@angular/http';
import { Observable } from 'rxjs/Observable';
import * as moment from 'moment';

import 'rxjs/add/operator/map';
import 'rxjs/add/operator/catch';
import 'rxjs/add/observable/throw';

import { environment } from '../../environments/environment';

import { User } from 'models/user';
import { Message } from 'models/message';
import { UserRegister } from 'models/user-register';
import { CurrentUser } from 'models/current-user';

@Injectable()
export class AuthService {
    //TODO: Add types to all methods!

    constructor(private http: Http) {
        this.fetchFromLocalStorage();
    }

    authenticatenow(user: User) {
        let variables = { login: user.login, password: user.password };
        let query = `mutation ($login: String $password: String)
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
    } &variables=` + JSON.stringify(variables);
        return this.http.get(environment.apiUrl + '/graphql?query=' + query)
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
    saveInLocalStorage(v: CurrentUser): void {
        localStorage.setItem('current_user', JSON.stringify(v));
        let expirDate = moment(JSON.parse(localStorage.getItem('current_user')).DataOfExpiration, 'YYYY-MM-DD hh:mm:ss[.sssssss]');
        let millisecondsBetweenExpirationAndNow: number = expirDate.milliseconds() - Date.now();
        if (v != null && millisecondsBetweenExpirationAndNow > 0) {
            this.setLogOutTimer(millisecondsBetweenExpirationAndNow);
        }
    }

    fetchFromLocalStorage(): void {
        let data = JSON.parse(localStorage.getItem('current_user'));
        if (data != null) {
            let expirDate = moment(JSON.parse(localStorage.getItem('current_user')).DataOfExpiration, 'YYYY-MM-DD hh:mm:ss[.sssssss]');
            let millisecondsBetweenExpirationAndNow: number = expirDate.valueOf() - Date.now();
            console.log(millisecondsBetweenExpirationAndNow);
            if (millisecondsBetweenExpirationAndNow < 0) {
                localStorage.removeItem('current_user');
            } else {
                this.setLogOutTimer(millisecondsBetweenExpirationAndNow);
            }
        }
    }
    setLogOutTimer(delayInMilliseconds: number): void {
        setTimeout(() => {
            localStorage.removeItem('current_user');
            window.location.reload(false);
        }, delayInMilliseconds);
    }

/*     checktoken() {
        console.log('*');
        if (localStorage.getItem('current_user')) {
            let today = moment();
            let other = moment(JSON.parse(localStorage.getItem('current_user')).DataOfExpiration, 'YYYY-MM-DD hh:mm:ss[.sssssss]');
            if (other.isBefore(today)) {
                localStorage.removeItem('current_user');
            } 
        }
    } */

    unauthentificatenow() {
        localStorage.removeItem('current_user');
    }

    validate(key_: string): Observable<Message> {
        let variable = { key: key_ };
        let query = `mutation ($key: String )
    {
      Auth
      {
        Activation(ActivationCode: $key)
         {
          Message
         }
      }
    } &variables=` + JSON.stringify(variable);
        return this.http.get(environment.apiUrl + '/graphql?query=' + query)
            .catch((error: any) => { return Observable.throw(error); }).map(res => {
                return res.json().data.Auth.Activation;
            });
    }

    register(user: UserRegister) {
        let variables = { login: user.login, nickname: user.nickname, password: user.password };
        let query = `mutation ($login: String $nickname: String $password: String)
    {
      Auth
      {
        Signup(Login: $login Nickname: $nickname Password: $password)
         {
           ActivationCode
         }
    }
    } &variables=` + JSON.stringify(variables);
        return this.http.get(environment.apiUrl + '/graphql?query=' + query)
            .catch((error: any) => { return Observable.throw(error); });
    }

}
