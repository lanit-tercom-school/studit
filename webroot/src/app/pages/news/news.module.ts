import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { NewsComponent } from './news.component';
import { NewsItemListComponent } from './news-list/news-list.component';

@NgModule({
  imports: [
    CommonModule
  ],
  exports: [
    NewsComponent,
  ],
  declarations: [NewsComponent, NewsItemListComponent]
})
export class NewsModule { }
