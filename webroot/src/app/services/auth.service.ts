import { Injectable } from '@angular/core';
import { User } from './../components/pages/authorization/user';
import { Http, Headers, Response } from '@angular/http';
import { Observable } from 'rxjs/Observable';
import 'rxjs/add/operator/map';
import 'rxjs/add/operator/catch';
import 'rxjs/add/observable/throw';

@Injectable()
export class AuthService {

    constructor(private http: Http) { }

    authenticatenow(user: User) {

        var headers = new Headers();

        headers.append('Content-Type', 'application/json');

        return this.http.post('http://localhost:8080/v1/auth/login/', JSON.stringify(user), { headers: headers })
            .map((res: Response) => {
                if (res.json().token) {
                    localStorage.setItem('auth_key', res.json().token);
                    localStorage.setItem('current_user', user.Login);
                }
                else
                    return Observable.throw('no token');
            })
            .catch((error: any) => { return Observable.throw(error) });
    }

    unauthentificatenow() {
        window.localStorage.removeItem("auth_key");
        return this.http.get('http://localhost:8080/v1/auth/logout/');
    }
}