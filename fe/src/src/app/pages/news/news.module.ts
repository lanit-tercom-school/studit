import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { NewsComponent } from './news.component';
import { NewsItemListComponent } from './news-item-list/news-item-list.component';

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
