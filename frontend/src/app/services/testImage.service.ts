import { Injectable } from '@angular/core';
import { Http, Headers, RequestOptions, Response } from '@angular/http';
import { Observable } from 'rxjs/Observable';


@Injectable()
export class TestImageService {

    constructor() {
    }
    testImage(URL: string, onError: ()=>void) {
        var tester = new Image();
        tester.onerror = onError;
        tester.src = URL;
    }
}