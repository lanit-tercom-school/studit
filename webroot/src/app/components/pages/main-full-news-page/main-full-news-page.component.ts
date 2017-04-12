import { Component, OnInit } from '@angular/core';

import {NewsItem} from "models/news-item";

@Component({
  selector: 'app-main-full-news-page',
  templateUrl: './main-full-news-page.component.html',
  styleUrls: ['./main-full-news-page.component.css']
})
export class MainFullNewsPageComponent implements OnInit {


newsItems = [
new NewsItem("https://upload.wikimedia.org/wikipedia/commons/thumb/9/96/Microsoft_logo_(2012).svg/2000px-Microsoft_logo_(2012).svg.png",
"26 декабря 2016",
"Новость НОМЕР",
"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam ac ex vitae lorem tristique pulvinar tempus rutrum erat. Nunc arcu mi, suscipit sit amet ullamcorper et, efficitur sit amet metus. Cras lectus nunc, aliquet id lacus vitae, dignissim aliquam ex. Nullam nec venenatis erat, eget tempor metus. Aenean tempus tempor mi sit amet pretium. Nam mollis mattis sodales. Nullam vitae tempor urna. Aenean elit elit, fringilla et viverra vel, convallis a eros. Donec tempus, justo sed fermentum mollis, neque erat consequat lorem, sed ornare sapien mi sit amet ex. Quisque ac augue condimentum justo molestie luctus. Aliquam erat volutpat.Suspendisse ut dui quis ligula vestibulum placerat. Pellentesque ac dui nec nunc ullamcorper sagittis. Ut leo libero, condimentum eget hendrerit et, vehicula eu tellus. Donec feugiat, lectus sit amet semper rutrum, libero ipsum elementum sapien, sit amet lobortis dui turpis quis dui. Suspendisse elit sem, egestas vel nibh eget, convallis venenatis nisi. Sed dui tellus, iaculis non semper vitae, placerat mattis eros. Nullam urna dolor, sagittis vel diam a, elementum porta libero. Etiam at pellentesque mi. In vestibulum accumsan euismod. Proin et vestibulum mauris, et consequat lacus. Aenean ut ornare arcu, nec molestie lacus. Integer lorem orci, fermentum eu ante sed, pulvinar pellentesque metus. Aliquam tellus ligula, ultricies nec interdum dapibus, consequat quis enim. Phasellus egestas nisi vitae orci iaculis, nec aliquet augue lobortis. Cras purus purus, scelerisque quis rutrum quis, congue eu urna. Suspendisse ornare blandit augue.")
];
myItems = this.newsItems[0];

  constructor() { }

  ngOnInit() {
  }

}
