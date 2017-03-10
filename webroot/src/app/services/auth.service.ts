import { Injectable } from '@angular/core';
import { User } from './../components/pages/authorization/user';
import { Http, Headers, Response } from '@angular/http';
import { Observable } from 'rxjs/Observable';
import 'rxjs/add/operator/map';

@Injectable()
export class AuthService {
    isAuthenticated: boolean = false;
    constructor(private http: Http) { }

    authenticatenow(user: User) {
        var headers = new Headers();

        headers.append('Content-Type', 'application/x-www-form-urlencoded');

        return this.http.post('http://localhost:8080/v1/auth/login/'
            , JSON.stringify({
                Login: user.email
                , Password: user.password
            })
            , { headers: headers }).map((res: Response) => {

                if (res.json().token) {
                    localStorage.setItem('auth_key', res.json().token);
                    return true;
                }
                else {
                    return false;
                }
            });
    }

    unauthentificatenow() {
        var headers = new Headers();

        headers.append('Content-Type', 'application/x-www-form-urlencoded');

        return new Promise((resolve) => {
            this.http.get('http://localhost:8080/v1/auth/logout/', { headers: headers });

            window.localStorage.removeItem("auth_key");
        })
    }
}