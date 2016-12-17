import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { HttpModule } from '@angular/http';

import { AppComponent } from './app.component';
import {NewsModule} from "./pages/news/news.module";
import {ProgressModule} from "./pages/progress/progress.module";

@NgModule({
  declarations: [
    AppComponent
  ],
  imports: [
    BrowserModule,
    FormsModule,
    HttpModule,
    NewsModule,
    ProgressModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
